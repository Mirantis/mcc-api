package lib

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"

	_ "k8s.io/client-go/plugin/pkg/client/auth/oidc"

	"golang.org/x/net/http/httpproxy"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"

	clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
	kaasv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
	apisutil "github.com/Mirantis/mcc-api/pkg/apis/util/common/v1alpha1"
	"github.com/Mirantis/mcc-api/pkg/client/internalclientset"
	"github.com/Mirantis/mcc-api/pkg/client/internalclientset/scheme"
	"github.com/Mirantis/mcc-api/pkg/errors"
	pkgutil "github.com/Mirantis/mcc-api/pkg/util"
)

const (
	ProviderTypeKey = "kaas.mirantis.com/provider"
	RegionKey       = "kaas.mirantis.com/region"

	retryIntervalResourceReady = 10 * time.Second
)

var (
	timeoutResourceDelete       = 15 * time.Minute
	retryIntervalResourceDelete = 10 * time.Second
)

var timeoutMachineReady = 30 * time.Minute
var ClusterGVK schema.GroupVersionKind

func init() {
	clusterGVKs, unversioned, err := scheme.Scheme.ObjectKinds(&clusterv1.Cluster{})
	if err != nil {
		panic(err)
	}
	if unversioned {
		panic(errors.Errorf("unversioned object %v", &clusterv1.Cluster{}))
	}
	ClusterGVK = clusterGVKs[0]
}

type Cluster struct {
	kubeClient     kubernetes.Interface
	clientSet      *internalclientset.Clientset
	kubeConfig     *rest.Config
	isWaitInfinite bool // for PollImmidiate  if need to wait under operator's control
}

type Interface interface {
	EnsureNamespace(namespaceName string) error
	setKubeClient() error
	GetOpenStackCredentialWithRetry(name, namespace string) (*kaasv1alpha1.OpenStackCredential, error)
	CreateOpenStackCredential(credential *kaasv1alpha1.OpenStackCredential) error
	CreatePublicKey(publicKey *kaasv1alpha1.PublicKey) error
	CreateClusterObject(cluster *clusterv1.Cluster) error
	CreateMachines(machines []*clusterv1.Machine, namespace string) error
	WaitForMachinesFold(interval, duration time.Duration, namespace, clusterName string, initialValue bool, foldfunc func(bool, *clusterv1.Machine, *kaasv1alpha1.MachineStatusMixin) bool) error
	GetMachinesForCluster(cluster *clusterv1.Cluster) ([]*clusterv1.Machine, error)
	PatchCluster(name, namespace string, data []byte) error
	DeleteCluster(namespace, name string) error
	waitForClusterDelete(namespace, name string) error
	GetCluster(namespace, name string) (*clusterv1.Cluster, error)
}

func NewCluster(kubeConfig *rest.Config, isWaitInfinite bool) (*Cluster, error) {
	kubeConfig.AuthConfigPersister = NewFakePersister()
	clientSet, err := internalclientset.NewForConfig(kubeConfig)
	if err != nil {
		return nil, errors.Wrap(err, "get clientset fails")
	}
	cluster := &Cluster{
		kubeConfig:     kubeConfig,
		clientSet:      clientSet,
		isWaitInfinite: isWaitInfinite,
	}
	if err := cluster.setKubeClient(); err != nil {
		return nil, err
	}
	return cluster, nil
}

func (c *Cluster) EnsureNamespace(namespaceName string) error {
	var namespaces *corev1.NamespaceList
	var err error
	if err = RetryOnError(c.isWaitInfinite, func() error {
		namespaces, err = c.kubeClient.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{
			FieldSelector: fmt.Sprintf("metadata.name=%s", namespaceName),
		})
		return err
	}); err != nil {
		return err
	}
	if len(namespaces.Items) > 0 {
		return nil
	}

	namespace := corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespaceName,
		},
	}
	err = RetryOnError(c.isWaitInfinite, func() error {
		_, err = c.kubeClient.CoreV1().Namespaces().Create(context.TODO(), &namespace, metav1.CreateOptions{})
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Cluster) setKubeClient() error {
	kclient, err := kubernetes.NewForConfig(c.kubeConfig)
	if err != nil {
		return errors.Wrap(err, "failed get Kubernetes client")
	}
	c.kubeClient = kclient
	return nil
}

func (c *Cluster) GetOpenStackCredentialWithRetry(name, namespace string) (*kaasv1alpha1.OpenStackCredential, error) {
	var res *kaasv1alpha1.OpenStackCredential
	var notFoundErr error
	err := RetryOnError(c.isWaitInfinite, func() error {
		var err error
		res, err = c.clientSet.KaasV1alpha1().OpenStackCredentials(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if k8serrors.IsNotFound(err) {
			notFoundErr = err
			return nil
		}
		return err
	})
	if notFoundErr != nil {
		return nil, notFoundErr
	}
	if err != nil {
		return nil, err
	}
	return res, err
}

func (c *Cluster) CreateOpenStackCredential(credential *kaasv1alpha1.OpenStackCredential) error {
	err := RetryOnError(c.isWaitInfinite, func() error {
		_, err := c.clientSet.KaasV1alpha1().OpenStackCredentials(credential.Namespace).Create(context.TODO(), credential, metav1.CreateOptions{})
		return err
	})
	return err
}

func (c *Cluster) CreatePublicKey(publicKey *kaasv1alpha1.PublicKey) error {
	var publicKeys *kaasv1alpha1.PublicKeyList
	var err error
	err = RetryOnError(c.isWaitInfinite, func() error {
		publicKeys, err = c.clientSet.KaasV1alpha1().PublicKeys(publicKey.Namespace).List(context.TODO(), metav1.ListOptions{
			FieldSelector: fmt.Sprintf("metadata.name=%s", publicKey.Name),
		})
		return err
	})
	if err != nil {
		return errors.Wrapf(err, "unable to list publicKeys in namespace %s", publicKey.Namespace)
	}
	if len(publicKeys.Items) > 0 {
		klog.Infof("publicKey %s already exists in namespace %s", publicKey.Name, publicKey.Namespace)
		return nil
	}
	err = RetryOnError(c.isWaitInfinite, func() error {
		_, err = c.clientSet.KaasV1alpha1().PublicKeys(publicKey.Namespace).Create(context.TODO(), publicKey, metav1.CreateOptions{})
		return err
	})
	if err != nil {
		return errors.Wrapf(err, "unable to create publicKey %s", publicKey.Name)
	}
	return nil
}

func (c *Cluster) CreateClusterObject(cluster *clusterv1.Cluster) error {
	namespace := metav1.NamespaceDefault
	if cluster.Namespace != "" {
		namespace = cluster.Namespace
	}

	var err error
	err = RetryOnError(c.isWaitInfinite, func() error {
		_, err = c.clientSet.ClusterV1alpha1().Clusters(namespace).Create(context.TODO(), cluster, metav1.CreateOptions{})
		return err
	})
	if err != nil {
		return errors.Wrapf(err, "error creating cluster in namespace %v", namespace)
	}
	return nil
}

func (c *Cluster) CreateMachines(machines []*clusterv1.Machine, namespace string) error {
	var (
		wg      sync.WaitGroup
		errOnce sync.Once
		gerr    error
	)
	// The approach to concurrency here comes from golang.org/x/sync/errgroup.
	for _, machine := range machines {
		wg.Add(1)

		go func(machine *clusterv1.Machine) {
			defer wg.Done()

			var createdMachine *clusterv1.Machine
			var err error
			err = RetryOnError(c.isWaitInfinite, func() error {
				createdMachine, err = c.clientSet.ClusterV1alpha1().Machines(namespace).Create(context.TODO(), machine, metav1.CreateOptions{})
				return err
			})
			if err != nil {
				errOnce.Do(func() {
					gerr = errors.Wrapf(err, "error creating a machine object in namespace %v", namespace)
				})
				return
			}

			if err = WaitForMachineReady(c.clientSet, createdMachine, c.isWaitInfinite); err != nil {
				errOnce.Do(func() { gerr = err })
			}
		}(machine)
	}
	wg.Wait()
	return gerr
}

func (c *Cluster) WaitForMachinesFold(interval, duration time.Duration, namespace, clusterName string, initialValue bool, foldfunc func(bool, *clusterv1.Machine, *kaasv1alpha1.MachineStatusMixin) bool) error {
	err := pkgutil.PollImmediate(c.isWaitInfinite, interval, duration, func() (bool, error) {
		machines, err := c.GetMachinesForCluster(&clusterv1.Cluster{
			TypeMeta: metav1.TypeMeta{
				Kind: "Cluster", // GetMachinesForCluster filters Machines by ownerReferences and checks Kind and Name fields
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: namespace,
				Name:      clusterName,
			},
		})
		if err != nil {
			klog.Infof("Failed to get Machines for cluster %s/%s: %v", namespace, clusterName, err)
			return false, nil
		}
		if len(machines) == 0 {
			return false, fmt.Errorf("no machines found for cluster %s/%s", namespace, clusterName)
		}
		result := initialValue
		for _, machine := range machines {
			status, err := apisutil.GetMachineStatus(machine)
			if err != nil {
				klog.Infof("Failed to parse status for Machine %s/%s: %v", machine.Namespace, machine.Name, err)
				return false, nil
			}
			result = foldfunc(result, machine, status)
		}
		return result, nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Cluster) GetMachinesForCluster(cluster *clusterv1.Cluster) ([]*clusterv1.Machine, error) {
	var machineslist []clusterv1.Machine
	var err error
	err = RetryOnError(c.isWaitInfinite, func() error {
		machineslist, err = GetMachinesForCluster(c.clientSet, cluster)
		return err
	})
	if err != nil {
		return nil, errors.Wrapf(err, "error listing Machines for Cluster %s/%s", cluster.Namespace, cluster.Name)
	}

	machines := make([]*clusterv1.Machine, len(machineslist))

	for i := range machineslist {
		machines[i] = &machineslist[i]
	}

	return machines, nil
}

func (c *Cluster) PatchCluster(name, namespace string, data []byte) error {
	var err error
	err = RetryOnError(c.isWaitInfinite, func() error {
		_, err = c.clientSet.ClusterV1alpha1().Clusters(namespace).Patch(context.TODO(), name, types.JSONPatchType, data, metav1.PatchOptions{})
		return err
	})
	if err != nil {
		return errors.Wrapf(err, "error patching cluster %s/%s", namespace, name)
	}
	return nil
}

// DeleteCluster deletes Cluster by name in a namespace
func (c *Cluster) DeleteCluster(namespace, name string) error {
	var err error
	err = RetryOnError(c.isWaitInfinite, func() error {
		err = c.clientSet.ClusterV1alpha1().Clusters(namespace).Delete(context.TODO(), name, newDeleteOptions())
		if !k8serrors.IsNotFound(err) {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.Wrapf(err, "error deleting Cluster %q in namespace %q", name, namespace)
	}
	err = c.waitForClusterDelete(namespace, name)
	if err != nil {
		return errors.Wrapf(err, "error waiting for Cluster %q deletion to complete in namespace %q", name, namespace)
	}
	return nil
}

func (c *Cluster) waitForClusterDelete(namespace, name string) error {
	return pkgutil.PollImmediate(c.isWaitInfinite, retryIntervalResourceDelete, timeoutResourceDelete, func() (bool, error) {
		klog.V(2).Infof("Waiting for Clusters to be deleted...")
		response, err := c.clientSet.ClusterV1alpha1().Clusters(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		if err != nil {
			if k8serrors.IsNotFound(err) {
				return true, nil
			}
			return false, nil
		}
		if response != nil {
			return false, nil
		}
		return true, nil
	})
}

func newDeleteOptions() metav1.DeleteOptions {
	propagationPolicy := metav1.DeletePropagationForeground
	return metav1.DeleteOptions{
		PropagationPolicy: &propagationPolicy,
	}
}

func WaitForMachineReady(cs internalclientset.Interface, machine *clusterv1.Machine, isWaitInfinite bool) error {
	timeout := timeoutMachineReady

	err := pkgutil.PollImmediate(isWaitInfinite, retryIntervalResourceReady, timeout, func() (bool, error) {
		klog.V(2).Infof("Waiting for Machine %v to become ready...", machine.Name)
		m, err := cs.ClusterV1alpha1().Machines(machine.Namespace).Get(context.TODO(), machine.Name, metav1.GetOptions{})
		if err != nil {
			return false, nil
		}

		// FIXME: This actually does nothing since we set UID as first annotation
		ready := m.Status.NodeRef != nil || len(m.Annotations) >= 1
		return ready, nil
	})

	return err
}

func ConnectToCluster(kubeconfigPath string, isWaitInfinite bool) (*Cluster, error) {
	kubeconfig, err := ioutil.ReadFile(kubeconfigPath)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read cluster kubeconfig")
	}
	restConfig, err := clientcmd.RESTConfigFromKubeConfig(kubeconfig)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse kubeconfig")
	}
	restConfig.Proxy = func(req *http.Request) (*url.URL, error) {
		return httpproxy.FromEnvironment().ProxyFunc()(req.URL)
	}
	var cluster *Cluster
	err = pkgutil.PollImmediate(isWaitInfinite, 10*time.Second, 5*time.Minute, func() (bool, error) {
		cluster, err = NewCluster(restConfig, isWaitInfinite)
		if err != nil {
			klog.Infof("failed to connect to cluster: %s, retrying...", err)
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to cluster")
	}
	return cluster, nil
}

func SetProviderLabel(meta *metav1.ObjectMeta, provider string) {
	if meta.Labels == nil {
		meta.Labels = make(map[string]string)
	}
	meta.Labels[ProviderTypeKey] = provider
}

func SetRegionLabel(meta *metav1.ObjectMeta, region string) {
	if meta.Labels == nil {
		meta.Labels = make(map[string]string)
	}
	meta.Labels[RegionKey] = region
}

func GetMachinesForCluster(cs internalclientset.Interface, cluster *clusterv1.Cluster) ([]clusterv1.Machine, error) {
	selector := labels.NewSelector()
	defReq, err := labels.NewRequirement(clusterv1.MachineClusterLabelName, selection.Equals, []string{cluster.Name})
	if err != nil {
		return nil, err
	}
	selector = selector.Add(*defReq)
	opts := metav1.ListOptions{
		LabelSelector: selector.String(),
	}
	machineList, err := cs.ClusterV1alpha1().Machines(cluster.Namespace).List(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	return machineList.Items, nil
}

func (c *Cluster) GetCluster(namespace, name string) (*clusterv1.Cluster, error) {
	var cluster *clusterv1.Cluster
	var err error
	err = RetryOnError(c.isWaitInfinite, func() error {
		cluster, err = c.clientSet.ClusterV1alpha1().Clusters(namespace).Get(context.TODO(), name, metav1.GetOptions{})
		return err
	})
	if err != nil {
		return nil, err
	}
	cluster.SetGroupVersionKind(ClusterGVK)
	return cluster, nil
}

package v1alpha1

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/klog"
	deploymentutil "k8s.io/kubectl/pkg/util/deployment"
	"sigs.k8s.io/controller-runtime/pkg/client"

	lcmv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/common/lcm/v1alpha1"
	clusterv1 "github.com/Mirantis/mcc-api/v2/pkg/apis/public/cluster/v1alpha1"
	"github.com/Mirantis/mcc-api/v2/pkg/errors"
)

// mandatoryReadinessCheckNamespaces list of namespaces resources to check for notReadyKaasObjects
var mandatoryReadinessCheckNamespaces = map[string]bool{"kube-system": true, "kaas": true}
var excludedReadinessCheckNamespaces = map[string]bool{"default": true}

func (s *ClusterStatusMixin) UpdateClusterState(ctx context.Context, cl, regionalClient client.Client, cluster *clusterv1.Cluster) error {
	klog.Infof("Updating cluster state %v/%v with the list of not ready resources", cluster.Namespace, cluster.Name)
	// getting namespace list to check the object readinness
	nsList, err := GetNamespacesList(ctx, regionalClient, cluster)
	if err != nil {
		return errors.Wrapf(err, "failed to get helmbundle related object namespaces for cluster %s/%s", cluster.Namespace, cluster.Name)
	}
	notReadyKaaSObjects, err := GetNotReadyObjects(cl, nsList)
	if err != nil {
		return errors.Wrapf(err, "failed to get not ready objects for cluster %s/%s", cluster.Namespace, cluster.Name)
	}
	s.NotReadyObjects = notReadyKaaSObjects
	return nil
}

// GetNotReadyObjects returns not readyobjects related to MCC
func GetNotReadyObjects(cl client.Client, nsReadinessList []string) (Objects, error) {
	var notReadyKaaSObjects Objects

	// List objects from interested namespaces
	for _, ns := range nsReadinessList {
		opts := &client.ListOptions{
			Namespace: ns,
		}
		services := &corev1.ServiceList{}
		err := cl.List(context.Background(), services, opts)
		if err != nil {
			return Objects{}, err
		}
		for _, service := range services.Items {
			if service.Spec.Type != corev1.ServiceTypeLoadBalancer {
				continue
			}
			if service.Status.LoadBalancer.Ingress == nil || (service.Status.LoadBalancer.Ingress[0].IP == "" && service.Status.LoadBalancer.Ingress[0].Hostname == "") {
				notReadyKaaSObjects.Services = append(notReadyKaaSObjects.Services, Service{Name: service.Name, Namespace: service.Namespace})
			}
		}

		deployments := &appsv1.DeploymentList{}
		err = cl.List(context.Background(), deployments, opts)
		if err != nil {
			return Objects{}, err
		}
		for _, deployment := range deployments.Items {
			ready, ctrl := isDeploymentReady(&deployment)
			if !ready {
				notReadyKaaSObjects.Deployments = append(notReadyKaaSObjects.Deployments, ctrl)
			}
		}

		statefulsets := &appsv1.StatefulSetList{}
		err = cl.List(context.Background(), statefulsets, opts)
		if err != nil {
			return Objects{}, err
		}
		for _, statefulset := range statefulsets.Items {
			ready, ctrl := isStatefulSetReady(&statefulset)
			if !ready {
				notReadyKaaSObjects.StatefulSets = append(notReadyKaaSObjects.StatefulSets, ctrl)
			}
		}

		daemonsets := &appsv1.DaemonSetList{}
		err = cl.List(context.Background(), daemonsets, opts)
		if err != nil {
			return Objects{}, err
		}
		for _, daemonset := range daemonsets.Items {
			ready, ds := isDaemonSetReady(&daemonset)
			if !ready {
				notReadyKaaSObjects.DaemonSets = append(notReadyKaaSObjects.DaemonSets, ds)
			}
		}
	}
	return notReadyKaaSObjects, nil
}

// based on https://github.com/kubernetes/kubernetes/blob/f610eee1/staging/src/k8s.io/kubectl/pkg/polymorphichelpers/rollout_status.go#L59
func isDeploymentReady(d *appsv1.Deployment) (bool, Controller) {
	ctrl := Controller{
		Name:          d.Name,
		Namespace:     d.Namespace,
		Replicas:      d.Status.Replicas,
		ReadyReplicas: d.Status.ReadyReplicas,
	}

	if d.Generation > d.Status.ObservedGeneration {
		return false, ctrl
	}
	cond := deploymentutil.GetDeploymentCondition(d.Status, appsv1.DeploymentProgressing)
	if cond != nil && cond.Reason == deploymentutil.TimedOutReason {
		return false, ctrl
	}
	if d.Spec.Replicas != nil && d.Status.UpdatedReplicas < *d.Spec.Replicas {
		return false, ctrl
	}
	if d.Status.Replicas > d.Status.UpdatedReplicas {
		return false, ctrl
	}
	if d.Status.AvailableReplicas < d.Status.UpdatedReplicas {
		return false, ctrl
	}

	return true, Controller{}
}

func isStatefulSetReady(sts *appsv1.StatefulSet) (bool, Controller) {
	ctrl := Controller{
		Name:          sts.Name,
		Namespace:     sts.Namespace,
		Replicas:      sts.Status.Replicas,
		ReadyReplicas: sts.Status.ReadyReplicas,
	}
	if sts.Generation > sts.Status.ObservedGeneration {
		return false, ctrl
	}
	if sts.Status.ReadyReplicas != sts.Status.Replicas {
		return false, ctrl
	}
	if sts.Status.UpdateRevision != sts.Status.CurrentRevision {
		return false, ctrl
	}
	return true, Controller{}
}

func isDaemonSetReady(ds *appsv1.DaemonSet) (bool, DaemonSet) {
	var err error
	var maxUnavailable int
	res := DaemonSet{
		Controller: Controller{
			Name:      ds.Name,
			Namespace: ds.Namespace,
		},
	}
	if ds.Spec.UpdateStrategy.Type == appsv1.RollingUpdateDaemonSetStrategyType {
		maxUnavailable, err = intstr.GetScaledValueFromIntOrPercent(ds.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable, int(ds.Status.DesiredNumberScheduled), true)
		if err != nil {
			maxUnavailable = int(ds.Status.DesiredNumberScheduled)
		}
		res.UpdatedNumberScheduled = ds.Status.UpdatedNumberScheduled
		res.DesiredNumberScheduled = ds.Status.DesiredNumberScheduled
	}
	expectedReady := int(ds.Status.DesiredNumberScheduled) - maxUnavailable
	res.Replicas = int32(expectedReady)
	res.ReadyReplicas = ds.Status.NumberReady

	if ds.Generation > ds.Status.ObservedGeneration {
		return false, res
	}
	if res.ReadyReplicas < res.Replicas || res.UpdatedNumberScheduled < res.DesiredNumberScheduled {
		return false, res
	}
	return true, DaemonSet{}
}

// GetNamespacesList returns list of namspaces from cluster related helmbudle. The mandatory and excluded values are took into account
func GetNamespacesList(ctx context.Context, cl client.Client, cluster *clusterv1.Cluster) ([]string, error) {
	namespaces := mandatoryReadinessCheckNamespaces
	// Get Helmbundle related namespace values
	helmBundle := &lcmv1alpha1.HelmBundle{}
	err := cl.Get(ctx, client.ObjectKey{Namespace: cluster.Namespace, Name: cluster.Name}, helmBundle)
	if err != nil {
		if k8serrors.IsNotFound(errors.Cause(err)) {
			return getNamespaces(namespaces), nil
		}
		return nil, errors.Wrap(err, "failed to get HelmBundle object")
	}
	for _, release := range helmBundle.Spec.Releases {
		if !excludedReadinessCheckNamespaces[release.Namespace] {
			namespaces[release.Namespace] = true
		}
	}
	return getNamespaces(namespaces), nil
}

func getNamespaces(nsMap map[string]bool) []string {
	var list []string
	for ns := range nsMap {
		list = append(list, ns)
	}
	return list
}

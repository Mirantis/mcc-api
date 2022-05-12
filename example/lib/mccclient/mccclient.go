package mccclient

import (
	"fmt"
	"net/url"
	"os"
	"time"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/klog/v2"

	exampleLib "github.com/Mirantis/mcc-api/example/lib"
	"github.com/Mirantis/mcc-api/example/lib/keycloak"
	openstackClient "github.com/Mirantis/mcc-api/example/lib/openstack"
	lcmv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/common/lcm/v1alpha1"
	clusterv1 "github.com/Mirantis/mcc-api/pkg/apis/public/cluster/v1alpha1"
	kaasv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
	apisutil "github.com/Mirantis/mcc-api/pkg/apis/util/common/v1alpha1"
	"github.com/Mirantis/mcc-api/pkg/errors"
	pkgutil "github.com/Mirantis/mcc-api/pkg/util"
	"github.com/gophercloud/utils/openstack/clientconfig"
)

const (
	retryKubeConfigReady   = 3 * time.Minute
	timeoutKubeconfigReady = 30 * time.Minute
	kubeAPIPort            = "443"
)

type MccClient struct {
	ManagementCluster *exampleLib.Cluster
	isWaitInfinite    bool
}

type Interface interface {
	CloudConfigExists(name, namespace string) (bool, error)
	CreateCloudConfig(region, name, namespace string, cloudConfig *clientconfig.Cloud) error
	CreateManagedCluster(cluster *clusterv1.Cluster, machines []*clusterv1.Machine, publicKey *kaasv1alpha1.PublicKey) error
	waitForMachinesReady(namespace, clusterName string) error
	waitForOneMachineNotReady(namespace, clusterName string) error
	GetKubeconfig(kubeconfigOutput string, clusterName, namespace, realm, username, password string) error
	Upgrade(clusterName, namespace, releaseName string) error
	Delete(clusterName, namespace string) error
}

func NewMccClient(managementClusterKubeconfig string) (*MccClient, error) {
	// Wait component to be rady overwrite approach
	var isWaitInfinite bool
	isWaitInfiniteFlag := os.Getenv("BOOTSTRAP_INFINITE_TIMEOUT")
	if isWaitInfiniteFlag == "true" {
		isWaitInfinite = true
	}
	klog.Infof("--- Conditions waiting while deploy process is %v ---", isWaitInfinite)

	managementCluster, err := exampleLib.ConnectToCluster(managementClusterKubeconfig, isWaitInfinite)
	if err != nil {
		return nil, errors.Wrap(err, "unable connect to management cluster")
	}
	return &MccClient{
		ManagementCluster: managementCluster,
		isWaitInfinite:    isWaitInfinite,
	}, nil
}

func (d *MccClient) CloudConfigExists(name, namespace string) (bool, error) {
	_, err := d.ManagementCluster.GetOpenStackCredentialWithRetry(name, namespace)
	if err != nil {
		if k8serrors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (d *MccClient) CreateCloudConfig(region, name, namespace string, cloudConfig *clientconfig.Cloud) error {
	err := d.ManagementCluster.EnsureNamespace(namespace)
	if err != nil {
		return errors.Wrapf(err, "unable to ensure namespace %q", namespace)
	}

	credential := openstackClient.CredentialFromCloud(cloudConfig, metav1.ObjectMeta{
		Name:      name,
		Namespace: namespace,
		Labels: map[string]string{
			"kaas.mirantis.com/region": region,
		},
	})
	klog.Infof("Creating OpenStackCredential in %s namespace", namespace)
	if err := d.ManagementCluster.CreateOpenStackCredential(credential); err != nil {
		return errors.Wrapf(err, "error creating OpenStackCredential %s in namespace %s", name, namespace)
	}
	return nil
}

func (d *MccClient) CreateManagedCluster(cluster *clusterv1.Cluster, machines []*clusterv1.Machine, publicKey *kaasv1alpha1.PublicKey) error {
	klog.Infof("Creating publicKey %v in namespace %v", publicKey.Name, publicKey.Namespace)
	err := d.ManagementCluster.CreatePublicKey(publicKey)
	if err != nil {
		return errors.Wrap(err, "failed to create publicKey")
	}

	klog.Infof("Creating cluster object %v in namespace %q", cluster.Name, cluster.Namespace)
	if err := d.ManagementCluster.CreateClusterObject(cluster); err != nil {
		return errors.Wrapf(err, "unable to create managed cluster %q in management cluster", cluster.Name)
	}
	klog.Infof("Creating machine objects for cluster %v in namespace %q", cluster.Name, cluster.Namespace)
	if err := d.ManagementCluster.CreateMachines(machines, cluster.Namespace); err != nil {
		return errors.Wrap(err, "unable to create machines")
	}

	err = d.waitForMachinesReady(cluster.Namespace, cluster.Name)
	if err != nil {
		return err
	}

	klog.Info("Done provisioning managed cluster.")
	return nil
}

func (d *MccClient) waitForMachinesReady(namespace, clusterName string) error {
	klog.Info("Waiting for cluster nodes to be Ready...")
	return d.ManagementCluster.WaitForMachinesFold(30*time.Second, 20*time.Minute, namespace, clusterName, true, func(res bool, machine *clusterv1.Machine, status *kaasv1alpha1.MachineStatusMixin) bool {
		if status.Status == lcmv1alpha1.LCMMachineStateReady {
			return res
		}
		klog.Infof("Machine %s/%s has status %s.", machine.Namespace, machine.Name, status.Status)
		return false
	})
}

func (d *MccClient) waitForOneMachineNotReady(namespace, clusterName string) error {
	klog.Info("Waiting for one node to become not Ready...")
	return d.ManagementCluster.WaitForMachinesFold(30*time.Second, 20*time.Minute, namespace, clusterName, false, func(res bool, machine *clusterv1.Machine, status *kaasv1alpha1.MachineStatusMixin) bool {
		if status.Status == lcmv1alpha1.LCMMachineStateReady {
			return res
		}
		klog.Infof("Machine %s/%s has status %s.", machine.Namespace, machine.Name, status.Status)
		return true
	})
}

func (d *MccClient) GetKubeconfig(kubeconfigOutput string, clusterName, namespace, realm, username, password string) error {
	err := WaitForOIDCReadiness(d.ManagementCluster, namespace, clusterName, d.isWaitInfinite)
	if err != nil {
		return err
	}
	klog.Info("Waiting for cluster kubeconfig to be ready...")
	err = pkgutil.PollImmediate(d.isWaitInfinite, retryKubeConfigReady, timeoutKubeconfigReady, func() (bool, error) {
		cluster, err := d.ManagementCluster.GetCluster(namespace, clusterName)
		if err != nil {
			return false, errors.Errorf("get cluster fails: %w", err)
		}
		err = generateKubeconfig(kubeconfigOutput, cluster, realm, username, password)
		if err != nil {
			klog.Infof("Managed cluster kubeconfig is not ready yet: %v, retrying...", err)
			return false, nil
		}
		return true, nil
	})

	if err != nil {
		return err
	}
	klog.Infof("Done generating kubeconfig. You can now access your cluster with kubectl --kubeconfig %v", kubeconfigOutput)
	return nil
}

func (d *MccClient) Upgrade(clusterName, namespace, releaseName string) error {
	klog.Info("Checking available releases")
	cluster, err := d.ManagementCluster.GetCluster(namespace, clusterName)
	if err != nil {
		return errors.Errorf("get cluster fails: %w", err)
	}
	status, err := apisutil.GetClusterStatus(cluster)
	if err != nil {
		return errors.Errorf("get cluster status fails: %w", err)
	}

	availableReleases := status.ReleaseRefs.Available
	if len(availableReleases) == 0 {
		klog.Info("Cluster is already latest version")
		return nil
	}

	if releaseName == "" {
		klog.Infof("No release name provided. Please specify --release-name. The list of available releases:\n")
		for _, release := range availableReleases {
			fmt.Println(release.Name)
		}
		return errors.Errorf("Unable to upgrade cluster %v in namespace %v", clusterName, namespace)
	}

	if status.ReleaseRefs.Current.Name == releaseName {
		klog.Infof("Current cluster release is already %s", releaseName)
		return nil
	}
	releaseAvailable := false
	for _, release := range availableReleases {
		if release.Name == releaseName {
			releaseAvailable = true
			break
		}
	}
	if !releaseAvailable {
		return errors.Errorf("Release %s is not available", releaseName)
	}

	klog.Infof("Upgrade managed cluster release to %v", releaseName)
	data := []byte(fmt.Sprintf(`[{"op":"replace","path":"/spec/providerSpec/value/release","value":"%s"}]`, releaseName))
	if err = d.ManagementCluster.PatchCluster(clusterName, namespace, data); err != nil {
		return err
	}

	// Wait for one of the machines to become not ready
	err = d.waitForOneMachineNotReady(namespace, clusterName)
	if err != nil {
		return err
	}

	// Wait for all machines to become ready
	err = d.waitForMachinesReady(namespace, clusterName)
	if err != nil {
		return err
	}

	klog.Infof("Cluster %s updated successfully. Current release version is %s", clusterName, releaseName)
	return nil
}

func (d *MccClient) Delete(clusterName, namespace string) error {
	cluster, err := d.ManagementCluster.GetCluster(namespace, clusterName)
	if err != nil {
		return errors.Errorf("get cluster fails: %w", err)
	}
	if cluster == nil {
		klog.Infof("Cluster %s does not exist in a namespace %s", clusterName, namespace)
		return nil
	}
	klog.Infof("Deleting cluster %q in namespace %q", clusterName, namespace)
	if err := d.ManagementCluster.DeleteCluster(namespace, clusterName); err != nil {
		return errors.Errorf("delete cluster fails: %w", err)
	}
	klog.Infof("Cluster %s removed successfully", clusterName)
	return nil
}

func WaitForOIDCReadiness(managementCluster *exampleLib.Cluster, namespace, name string, isWaitInfinite bool) error {
	klog.Info("Waiting for OIDC configuration readiness...")
	err := pkgutil.PollImmediate(isWaitInfinite, 10*time.Second, 20*time.Minute, func() (bool, error) {
		cluster, err := managementCluster.GetCluster(namespace, name)
		if err != nil {
			klog.Infof("failed to get cluster %s/%s, retrying", namespace, name)
			return false, nil
		}
		status, err := apisutil.GetClusterStatus(cluster)
		if err != nil {
			return false, errors.Wrap(err, "failed to parse cluster status")
		}
		if status.OIDC == nil || !status.OIDC.Ready {
			klog.Info("OIDC configuration is not ready yet, retrying")
			return false, nil
		}
		return true, nil
	})
	if err != nil {
		return errors.Wrap(err, "failed waiting for OIDC configuration readiness")
	}
	return nil
}

func generateKubeconfig(kubeconfigOutput string, cluster *clusterv1.Cluster, realm, username, password string) error {
	status, err := apisutil.GetClusterStatus(cluster)
	if err != nil {
		return errors.Errorf("get cluster status fails: %w", err)
	}

	if status.APIServerCertificate == nil {
		return errors.New("failed to get api server certificate")
	}

	if status.OIDC == nil {
		return errors.New("failed to get OIDC config from cluster status")
	}
	keycloakURL, err := url.Parse(status.OIDC.IssuerURL)
	if err != nil {
		return errors.Errorf("invalid issuer URL: %w", err)
	}

	keycloakConfig := keycloak.Config{
		BasePath:  fmt.Sprintf("%s://%s", keycloakURL.Scheme, keycloakURL.Host),
		Username:  username,
		Password:  password,
		AuthRealm: realm,
	}

	keycloakClient, err := keycloak.NewKeycloak(keycloakConfig)
	if err != nil {
		return errors.Errorf("unable to create keycloak client: %w", err)
	}

	token, err := keycloakClient.GetToken(status.OIDC.ClientID)
	if err != nil {
		return errors.Errorf("unable to get keycloak token: %w", err)
	}

	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	config, err := loadingRules.GetStartingConfig()
	if err != nil {
		return err
	}

	config.Clusters[cluster.GetName()] = &clientcmdapi.Cluster{
		CertificateAuthorityData: status.APIServerCertificate,
		Server:                   fmt.Sprintf("https://%s:%s", status.LoadBalancerHost, kubeAPIPort),
	}
	contextName := fmt.Sprintf("%s@%s@%s", username, cluster.GetNamespace(), cluster.GetName())
	config.Contexts[contextName] = &clientcmdapi.Context{
		Cluster:  cluster.GetName(),
		AuthInfo: username,
	}
	config.CurrentContext = contextName
	config.AuthInfos[username] = &clientcmdapi.AuthInfo{
		AuthProvider: &clientcmdapi.AuthProviderConfig{
			Name: "oidc",
			Config: map[string]string{
				"client-id":                      status.OIDC.ClientID,
				"idp-certificate-authority-data": status.OIDC.Certificate,
				"idp-issuer-url":                 status.OIDC.IssuerURL,
				"refresh-token":                  token.RefreshToken,
				"id-token":                       token.IDToken,
				"token":                          token.AccessToken,
			},
		},
	}
	err = clientcmd.WriteToFile(*config, kubeconfigOutput)
	if err != nil {
		return errors.Errorf("unable to write kubeconfig: %w", err)
	}
	return nil
}

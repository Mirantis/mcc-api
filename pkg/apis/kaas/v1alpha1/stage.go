package v1alpha1

import (
	lcmv1 "github.com/Mirantis/mcc-api/v2/pkg/apis/lcm/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// +gocode:public-api=true
	NetworkPreparation lcmv1.StageName = "Network prepared"
	// +gocode:public-api=true
	LoadBalancersCreation lcmv1.StageName = "Load balancers created"
	// +gocode:public-api=true
	IAMScopesCreation lcmv1.StageName = "IAM objects created"
	// +gocode:public-api=true
	BastionCreation lcmv1.StageName = "Bastion created"
	// +gocode:public-api=true
	OIDCBecomingReady lcmv1.StageName = "OIDC configured"
	// +gocode:public-api=true
	KubernetesAPIStarting lcmv1.StageName = "Kubernetes API started"
	// +gocode:public-api=true
	CephClusterCreation lcmv1.StageName = "Ceph cluster created"
	// +gocode:public-api=true
	HelmControllerDeployment lcmv1.StageName = "Helm-controller deployed"
	// +gocode:public-api=true
	HelmBundleCreation lcmv1.StageName = "HelmBundle created"
	// +gocode:public-api=true
	CertificatesCreation lcmv1.StageName = "Certificates configured"
	// +gocode:public-api=true
	MachinesBecomingReady lcmv1.StageName = "All machines of the cluster are ready"
	// +gocode:public-api=true
	ClusterBecomingReady lcmv1.StageName = "Cluster is ready"

	// +gocode:public-api=true
	ProviderInstanceReady lcmv1.StageName = "Provider instance ready"
)

// LCMStages --
// +gocode:public-api=true
var LCMStages = []lcmv1.StageName{
	KubernetesAPIStarting,
	HelmControllerDeployment,
	HelmBundleCreation,
	CertificatesCreation,
	MachinesBecomingReady,
	OIDCBecomingReady,
	ClusterBecomingReady,
}

// ClusterDeploymentStatusList contains a list of ClusterDeploymentStatus
// +gocode:public-api=true
type ClusterDeploymentStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ClusterDeploymentStatus `json:"items"`
}

// ClusterDeploymentStatus describes cluster deployment process
// +k8s:openapi-gen=true
// +kubebuilder:resource
type ClusterDeploymentStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Release string          `json:"release,omitempty"`
	Stages  lcmv1.StageList `json:"stages"`
}

// MachineDeploymentStatus describes machine deployment process
// +k8s:openapi-gen=true
// +kubebuilder:resource
type MachineDeploymentStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Release string          `json:"release"`
	Stages  lcmv1.StageList `json:"stages"`
}

// ClusterUpgradeStatusList contains a list of ClusterUpgradeStatus
// +gocode:public-api=true
type ClusterUpgradeStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ClusterUpgradeStatus `json:"items"`
}

// ClusterUpgradeStatus describes cluster upgrade process
// +k8s:openapi-gen=true
// +kubebuilder:resource
type ClusterUpgradeStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	FromRelease string          `json:"fromRelease"`
	ToRelease   string          `json:"toRelease"`
	Stages      lcmv1.StageList `json:"stages"`
}

// MachineUpgradeStatusList contains a list of MachineUpgradeStatus
// +gocode:public-api=true
type MachineUpgradeStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MachineUpgradeStatus `json:"items"`
}

// MachineDeploymentStatusList contains a list of MachineDeploymentStatus
// +gocode:public-api=true
type MachineDeploymentStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MachineDeploymentStatus `json:"items"`
}

// MachineUpgradeStatus describes machine upgrade process
// +k8s:openapi-gen=true
// +kubebuilder:resource
type MachineUpgradeStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	FromRelease string          `json:"fromRelease"`
	ToRelease   string          `json:"toRelease"`
	Stages      lcmv1.StageList `json:"stages"`
}

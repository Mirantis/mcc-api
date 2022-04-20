package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kaasmiracephv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/common/kaascephcluster/v1alpha1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KaaSCephCluster is the Schema for the kaascephclusters API
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +kubebuilder:subresource:status
type KaaSCephCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KaaSCephClusterSpec                     `json:"spec"`
	Status *kaasmiracephv1alpha1.CephClusterStatus `json:"status,omitempty"`
}

type KaaSCephClusterSpec struct {
	K8sCluster *v1.ObjectReference `json:"k8sCluster"`

	CephClusterSpec *kaasmiracephv1alpha1.CephClusterSpec `json:"cephClusterSpec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KaaSCephClusterList contains a list of KaaSCephCluster
type KaaSCephClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []KaaSCephCluster `json:"items"`
}

// KaaSCephOperationRequest is the Schema for the kaascephoperationrequests API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +kubebuilder:subresource:status
type KaaSCephOperationRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KaaSCephOperationRequestSpec                     `json:"spec"`
	Status *kaasmiracephv1alpha1.CephOperationRequestStatus `json:"status,omitempty"`
}

type KaaSCephOperationRequestSpec struct {
	KaaSCephCluster *v1.ObjectReference `json:"kaasCephCluster"`

	kaasmiracephv1alpha1.CephOperationRequestSpec `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// KaaSCephOperationRequestList contains a list of KaaSCephOperationRequest
type KaaSCephOperationRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []KaaSCephOperationRequest `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KaaSCephCluster{}, &KaaSCephClusterList{}, &KaaSCephOperationRequest{}, &KaaSCephOperationRequestList{})
}

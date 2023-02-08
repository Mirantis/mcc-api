package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kaasv1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/public/kaas/v1alpha1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeStorage defines drives specification on the Node
// +k8s:openapi-gen=true
// +kubebuilder:resource:scope=Cluster
type NodeStorage struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status NodeStorageStatus `json:"status,omitempty"`
}

// NodeStorageStatus contains information about storage
type NodeStorageStatus struct {
	// Drives is the list of detected Node drives
	Drives []*kaasv1alpha1.MachineStorage `json:"drives,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeStorageList is a list of NodeStorage resources
type NodeStorageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []NodeStorage `json:"items"`
}

func init() {
	SchemeBuilder.Register(&NodeStorage{}, &NodeStorageList{})
}

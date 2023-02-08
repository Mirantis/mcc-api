package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MachinePollStatus describes machine conditions that are updated as a result of polling
// +k8s:openapi-gen=true
// +gocode:public-api=true
type MachinePollStatus struct {
	metav1.TypeMeta       `json:",inline"`
	metav1.ObjectMeta     `json:"metadata,omitempty"`
	ProviderInstanceState InstanceState `json:"instanceState,omitempty"`
	ConditionsSummary     `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MachinePollStatusList contains a list of MachinePollStatus
// +gocode:public-api=true
type MachinePollStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MachinePollStatus `json:"items"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&MachinePollStatus{}, &MachinePollStatusList{})
}

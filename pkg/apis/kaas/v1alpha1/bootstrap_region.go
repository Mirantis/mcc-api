package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BootstrapRegion is the Schema for the Bootstrap Region API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Provider",type=string,JSONPath=`.spec.provider`,description="Provider Type"
// +kubebuilder:printcolumn:name="Ready",type=string,JSONPath=`.status.ready`,description="Ready"
// +kubebuilder:printcolumn:name="Age",type=date,JSONPath=`.metadata.creationTimestamp`
// +gocode:public-api=true
type BootstrapRegion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BootstrapRegionSpec   `json:"spec"`
	Status BootstrapRegionStatus `json:"status,omitempty"`
}

func (br *BootstrapRegion) InitConditions() {
	br.Status.Conditions = []Condition{
		{Type: HelmCondition, Message: "Bootstrap helm bundle is not ready", Ready: false},
		{Type: ProviderCondition, Message: "Provider charts are not configured in bundle", Ready: false},
		{Type: DeploymentsCondition, Message: "Provider deployments are not ready", Ready: false},
	}
}
func (br *BootstrapRegion) UpdateCondition(cType ConditionType, message string, ready bool) {
	c := GetConditionFromSummary(br.Status.ConditionsSummary, cType)
	if c != nil {
		c.Message = message
		c.Ready = ready
	} else {
		br.Status.Conditions = append(br.Status.Conditions,
			Condition{Type: cType, Message: message, Ready: ready})
	}
	br.Status.ConditionsSummary = GetConditionsSummary(br.Status.Conditions)
}

type BootstrapRegionStatus struct {
	ConditionsSummary `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// BootstrapRegionList contains a list of BootstrapRegion
// +gocode:public-api=true
type BootstrapRegionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []BootstrapRegion `json:"items"`
}
type BootstrapRegionSpec struct {
	Provider string `json:"provider"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&BootstrapRegion{}, &BootstrapRegionList{})
}

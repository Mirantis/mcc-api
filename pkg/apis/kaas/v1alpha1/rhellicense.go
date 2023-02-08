package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RHELLicenseList contains a list of RHELLicense
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type RHELLicenseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []RHELLicense `json:"items"`
}
type RHELLicenseSpec struct {
	Username string       `json:"username,omitempty"`
	Password *SecretValue `json:"password,omitempty"`

	OrgID         string       `json:"orgID,omitempty"`
	ActivationKey *SecretValue `json:"activationKey,omitempty"`

	PoolIDs []string `json:"poolIDs,omitempty"`
	RpmURL  string   `json:"rpmUrl,omitempty"`
}

// RHELLicense is the Schema for the Red Hat license API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:resource
// +gocode:public-api=true
type RHELLicense struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RHELLicenseSpec `json:"spec,omitempty"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&RHELLicense{}, &RHELLicenseList{})
}

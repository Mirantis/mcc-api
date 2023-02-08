package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type LicenseStatus struct {
	Instance        string         `json:"instance,omitempty"`
	CustomerID      string         `json:"customerID,omitempty"`
	Dev             bool           `json:"dev,omitempty"`
	Limits          *LicenseLimits `json:"limits,omitempty"`
	OpenstackLimits *LicenseLimits `json:"openstack,omitempty"`
	ExpirationTime  metav1.Time    `json:"expirationTime,omitempty"`
	Expired         bool           `json:"expired"`
}

// License is the Schema for the Mirantis Container Cloud license API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +kubebuilder:resource:scope=Cluster
// +gocode:public-api=true
type License struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LicenseSpec   `json:"spec,omitempty"`
	Status LicenseStatus `json:"status,omitempty"`
}

// LicenseList contains a list of License
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type LicenseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []License `json:"items"`
}
type LicenseSpec struct {
	License *SecretValue `json:"license,omitempty"`
}
type LicenseLimits struct {
	Clusters          int `json:"clusters"`
	WorkersPerCluster int `json:"workersPerCluster"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&License{}, &LicenseList{})
}

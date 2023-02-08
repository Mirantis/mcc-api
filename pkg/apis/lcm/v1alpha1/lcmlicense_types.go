package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RHELLicenseSubscriptionList contains a list of RHELLicenseSubscription objects
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type RHELLicenseSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RHELLicenseSubscription `json:"items"`
}

// LicenseSubscriptionStatus defines the observed state of RHELLicenseSubscription
type LicenseSubscriptionStatus struct {
	// true when license was applied, otherwise false
	Applied bool `json:"applied,omitempty"`
	// error message when not applied
	Error      string `json:"error,omitempty"`
	UUID       string `json:"uuid,omitempty"`
	SystemName string `json:"systemName,omitempty"`
	// LastAttemptTime holds the last delete attempt timestamp
	LastAttemptTime metav1.Time `json:"lastAttemptTime,omitempty"`
}

// LicenseSecret -
type LicenseSecret struct {
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
}

// LicenseSecretItem -
type LicenseSecretItem struct {
	Secret LicenseSecret `json:"secret,omitempty"`
}

// RHELLicenseSubscription the LCMMachine corresponded license object
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="type",type="string",JSONPath=".spec.type",description="Type",priority=0
// +gocode:public-api=true
type RHELLicenseSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LicenseSubscriptionSpec   `json:"spec,omitempty"`
	Status LicenseSubscriptionStatus `json:"status,omitempty"`
}

// LicenseSubscriptionSpec defines the desired state of RHELLicenseSubscription
type LicenseSubscriptionSpec struct {
	// Username fields holds license username value
	Username string `json:"username,omitempty"`
	// PoolIDs fields holds license related poolIDs list
	PoolIDs []string `json:"poolIDs,omitempty"`
	// Password fields holds the password secret object reference
	Password LicenseSecretItem `json:"password,omitempty"`
	// ActivationKey fields holds the activation key secret object reference
	ActivationKey LicenseSecretItem `json:"activationKey,omitempty"`
	// ProxySecret holds the proxy related secret name
	ProxySecret string `json:"proxySecret,omitempty"`
	// Organisation ID which owns subscription
	OrgID string `json:"orgID,omitempty"`
	// Satellite Rpm for subscription manager
	RpmURL string `json:"rpmUrl,omitempty"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&RHELLicenseSubscription{}, &RHELLicenseSubscriptionList{})
}

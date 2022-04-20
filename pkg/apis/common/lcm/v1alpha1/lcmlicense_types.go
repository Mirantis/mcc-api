/*
Copyright 2020 The Mirantis Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RHELLicenseSubscription the LCMMachine corresponded license object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="type",type="string",JSONPath=".spec.type",description="Type",priority=0
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

// LicenseSecretItem -
type LicenseSecretItem struct {
	Secret LicenseSecret `json:"secret,omitempty"`
}

// LicenseSecret -
type LicenseSecret struct {
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RHELLicenseSubscriptionList contains a list of RHELLicenseSubscription objects
type RHELLicenseSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RHELLicenseSubscription `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RHELLicenseSubscription{}, &RHELLicenseSubscriptionList{})
}

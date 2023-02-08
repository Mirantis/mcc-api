package v1alpha1

// +gocode:public-api=true
type IAMStatus struct {
	Keycloak IAMComponentStatus `json:"keycloak,omitempty"`
}
type IAMComponentStatus struct {
	ComponentStatus `json:",inline"`
}

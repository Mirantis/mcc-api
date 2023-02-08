package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ScopeSpec defines the desired state of Scope
// +gocode:public-api=true
type ScopeSpec struct {
	ScopeFQN              string                   `json:"scopeFQN,omitempty"`
	Creator               string                   `json:"creator,omitempty"`
	Roles                 []*RoleDetails           `json:"roles,omitempty"`
	GrantInheritanceRules []*GrantsInheritanceRule `json:"grantInheritanceRules,omitempty"`
	OIDCRequired          bool                     `json:"oidcRequired,omitempty"`
	OIDCConfig            *OIDCConfig              `json:"oidcConfig,omitempty"`
	Description           string                   `json:"description,omitempty"`
}

// Scope represents existing scopes
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +k8s:openapi-gen=true
// +gocode:public-api=true
type Scope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScopeSpec   `json:"spec,omitempty"`
	Status ScopeStatus `json:"status,omitempty"`
}

// Role model
type RoleDetails struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

const (
	// +gocode:public-api=true
	ScopeFinalizer = "scope.iam.mirantis.com"
	// +gocode:public-api=true
	ClusterRedirectFinalizer = "clusterredirect.iam.mirantis.com"
)

// ScopeList contains a list of Scope
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type ScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Scope `json:"items"`
}

// Describes inheritence of grants for this scope
type GrantsInheritanceRule struct {
	GrantFQN       string `json:"grantFQN,omitempty"`
	ParentGrantFQN string `json:"parentGrantFQN,omitempty"`
}

// ScopeStatus defines the observed state of Scope
type ScopeStatus struct {
	// ClusterOIDC stores OIDC data for cluster scope
	ClusterOIDC *OIDCConfig `json:"oidcConfig,omitempty"`
}

// OIDC config model
type OIDCConfig struct {
	IssuerURL    string `json:"issuer_url,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	GroupClaim   string `json:"group_claim,omitempty"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&Scope{}, &ScopeList{})
}

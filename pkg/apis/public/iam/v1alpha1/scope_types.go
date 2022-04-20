/*
Copyright 2022 The Kubernetes Authors.

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

const (
	ScopeFinalizer = "scope.iam.mirantis.com"
)

// ScopeSpec defines the desired state of Scope
type ScopeSpec struct {
	ScopeFQN              string                   `json:"scopeFQN,omitempty"`
	Creator               string                   `json:"creator,omitempty"`
	Roles                 []*RoleDetails           `json:"roles,omitempty"`
	GrantInheritanceRules []*GrantsInheritanceRule `json:"grantInheritanceRules,omitempty"`
	OIDCRequired          bool                     `json:"oidcRequired,omitempty"`
	OIDCConfig            *OIDCConfig              `json:"oidcConfig,omitempty"`
	Description           string                   `json:"description,omitempty"`
	RedirectURIs          []string                 `json:"redirectURIs,omitempty"`
}

// ScopeStatus defines the observed state of Scope
type ScopeStatus struct {
	// ClusterOIDC stores OIDC data for cluster scope
	ClusterOIDC *OIDCConfig `json:"oidcConfig,omitempty"`
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster

// Scope represents existing scopes
// +k8s:openapi-gen=true
type Scope struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScopeSpec   `json:"spec,omitempty"`
	Status ScopeStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ScopeList contains a list of Scope
type ScopeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Scope `json:"items"`
}

// Role model
type RoleDetails struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

// Describes inheritence of grants for this scope
type GrantsInheritanceRule struct {
	GrantFQN       string `json:"grantFQN,omitempty"`
	ParentGrantFQN string `json:"parentGrantFQN,omitempty"`
}

// OIDC config model
type OIDCConfig struct {
	IssuerURL    string `json:"issuer_url,omitempty"`
	ClientID     string `json:"client_id,omitempty"`
	ClientSecret string `json:"client_secret,omitempty"`
	GroupClaim   string `json:"group_claim,omitempty"`
}

func init() {
	SchemeBuilder.Register(&Scope{}, &ScopeList{})
}

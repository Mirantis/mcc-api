package v1alpha1

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// +gocode:public-api=true
	GlobalRoleScope RoleScope = "global"
	// +gocode:public-api=true
	NamespaceRoleScope RoleScope = "namespace"
	// +gocode:public-api=true
	ClusterRoleScope RoleScope = "cluster"
	// Unknown is used only for checks
	// +gocode:public-api=true
	UnknownRoleScope RoleScope = "unknown"
)

// +gocode:public-api=true
type RoleScope string

// IAMRole represents the default roles
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:scope=Cluster
// +gocode:public-api=true
type IAMRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Scope RoleScope `json:"scope"`
	// +optional
	Description string `json:"description,omitempty"`
}

func (role IAMRole) ConstructKeycloakScopes(namespacesOrNamespacedClusters []string) []string {
	var keycloakScopes []string
	switch role.Scope {
	case ClusterRoleScope, NamespaceRoleScope:
		for _, nsOrCluster := range namespacesOrNamespacedClusters {
			keycloakScopes = append(keycloakScopes, fmt.Sprintf("m:%s:%s@%s", RoleKeycloakCliendIDMap[role.Name], nsOrCluster, role.Name))
		}
	case GlobalRoleScope:
		keycloakScopes = append(keycloakScopes, fmt.Sprintf("m:%s@%s", RoleKeycloakCliendIDMap[role.Name], role.Name))
	}
	return keycloakScopes
}

// RoleKeycloakCliendIDMap contains key-value pairs,
// where:
// - key is role
// - value is client ID.
// +gocode:public-api=true
var RoleKeycloakCliendIDMap = map[string]string{
	"stacklight-admin": "sl",
	"cluster-admin":    "k8s",
	"user":             "kaas",
	"operator":         "kaas",
	"bm-pool-operator": "kaas",
	"member":           "kaas",
	"global-admin":     "kaas",
}

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type IAMRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []IAMRole `json:"items"`
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(
		&IAMRole{},
		&IAMRoleList{},
	)
}

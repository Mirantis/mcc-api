package v1alpha1

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	GlobalRoleScope    RoleScope = "global"
	NamespaceRoleScope RoleScope = "namespace"
	ClusterRoleScope   RoleScope = "cluster"
	// Unknown is used only for checks
	UnknownRoleScope RoleScope = "unknown"
)

var RoleKeycloakCliendIDMap = map[string]string{
	"stacklight-admin": "sl",
	"cluster-admin":    "k8s",
	"user":             "kaas",
	"operator":         "kaas",
	"bm-pool-operator": "kaas",
	"global-admin":     "kaas",
}

type RoleScope string

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IAMRole represents the default roles
// +kubebuilder:resource:scope=Cluster
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

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type IAMRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []IAMRole `json:"items"`
}

func init() {
	SchemeBuilder.Register(
		&IAMRole{},
		&IAMRoleList{},
	)
}

package v1alpha1

import (
	"github.com/Mirantis/mcc-api/v2/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterRedirectSpec defines the desired state of ClusterRedirect
type ClusterRedirectSpec struct {
	UI         Redirects `json:"ui,omitempty"`
	Lens       Redirects `json:"lens,omitempty"`
	UCP        Redirects `json:"ucp,omitempty"`
	Stacklight Redirects `json:"stacklight,omitempty"`
	Openstack  Redirects `json:"openstack,omitempty"`
}

// ClusterRedirectStatus defines the observed state of ClusterRedirect
type ClusterRedirectStatus struct{}

// ClusterRedirect represents existing cluster redirects in Keycloak
// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type ClusterRedirect struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterRedirectSpec   `json:"spec,omitempty"`
	Status ClusterRedirectStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ClusterRedirectList contains a list of ClusterRedirect
// +gocode:public-api=true
type ClusterRedirectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterRedirect `json:"items"`
}

// +gocode:public-api=true
type Redirect struct {
	ClientID string   `json:"clientID,omitempty"`
	URLs     []string `json:"urls,omitempty"`
}

// +gocode:public-api=true
type Redirects []Redirect

func (re Redirects) AddToMap(m map[string][]string) {
	for _, r := range re {
		if len(r.URLs) != 0 {
			m[r.ClientID] = append(m[r.ClientID], r.URLs...)
		}
	}
}
func (re *Redirects) Merge(newRedirects Redirects) {
	for i := range newRedirects {
		if len(newRedirects[i].URLs) != 0 {
			dereferenced := *re
			oldRedirIndex := getIndexByClientID(dereferenced, newRedirects[i].ClientID)
			if oldRedirIndex != -1 {
				diffRedirs := util.Diff(dereferenced[oldRedirIndex].URLs, newRedirects[i].URLs)
				(*re)[oldRedirIndex].URLs = append(dereferenced[oldRedirIndex].URLs, diffRedirs...)
				continue
			}
			*re = append(dereferenced, Redirect{
				ClientID: newRedirects[i].ClientID,
				URLs:     newRedirects[i].URLs,
			})
		}
	}
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&ClusterRedirect{}, &ClusterRedirectList{})
}

// getIndexByClientID returns an index of Redirect with specified
// clientID
// +gocode:public-api=true
func getIndexByClientID(redirects Redirects, clientID string) int {
	for i, redirect := range redirects {
		if redirect.ClientID == clientID {
			return i
		}
	}
	return -1
}

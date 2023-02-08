package v1alpha1

import (
	"github.com/Masterminds/semver"
	"github.com/Mirantis/mcc-api/v2/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sort"
)

// +gocode:public-api=true
type indexKaasReleaseSort struct {
	target []IndexKaaSRelease
	err    error
}

func (p *indexKaasReleaseSort) Len() int {
	return len(p.target)
}
func (p *indexKaasReleaseSort) Less(l, r int) bool {
	lv, err := semver.NewVersion(p.target[l].Version)
	if err != nil {
		p.err = errors.Wrapf(err, "failed to parse release version '%s'", p.target[l].Version)
		return false
	}
	rv, err := semver.NewVersion(p.target[r].Version)
	if err != nil {
		p.err = errors.Wrapf(err, "failed to parse release version '%s'", p.target[r].Version)
		return false
	}
	return lv.Compare(rv) < 0
}
func (p *indexKaasReleaseSort) Swap(i, j int) {
	p.target[i], p.target[j] = p.target[j], p.target[i]
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +gocode:public-api=true
type Index struct {
	metav1.TypeMeta `json:",inline"`
	KaaSReleases    []IndexKaaSRelease `json:"kaasReleases"`
}
type IndexKaaSRelease struct {
	Version string `json:"version"`
	Path    string `json:"path"`
	SHA256  string `json:"sha256"`
}

// +gocode:public-api=true
func NewIndex() *Index {
	return &Index{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "kaas.mirantis.com/v1alpha1",
			Kind:       "Index",
		},
	}
}

// +gocode:public-api=true
func SortIndexKaasReleases(s []IndexKaaSRelease) error {
	si := indexKaasReleaseSort{target: s}
	sort.Sort(&si)
	if si.err != nil {
		return si.err
	}
	return nil
}

// +gocode:public-api=true
func init() {
	SchemeBuilder.Register(&Index{})
}

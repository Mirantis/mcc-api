package v1alpha1

import (
	"sort"

	"github.com/Masterminds/semver"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Mirantis/mcc-api/pkg/errors"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Index struct {
	metav1.TypeMeta `json:",inline"`
	KaaSReleases    []IndexKaaSRelease `json:"kaasReleases"`
}

type IndexKaaSRelease struct {
	Version string `json:"version"`
	Path    string `json:"path"`
	SHA256  string `json:"sha256"`
}

func NewIndex() *Index {
	return &Index{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "kaas.mirantis.com/v1alpha1",
			Kind:       "Index",
		},
	}
}

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

func SortIndexKaasReleases(s []IndexKaaSRelease) error {
	si := indexKaasReleaseSort{target: s}
	sort.Sort(&si)
	if si.err != nil {
		return si.err
	}
	return nil
}

func init() {
	SchemeBuilder.Register(&Index{})
}

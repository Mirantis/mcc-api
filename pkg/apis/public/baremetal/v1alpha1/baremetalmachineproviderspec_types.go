/*
Copyright 2019 The Kubernetes Authors.

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
	"net/url"
	"os"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/selection"

	metal3v1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/common/baremetal-operator/metal3.io/v1alpha1"
	kaasv1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
	"github.com/Mirantis/mcc-api/pkg/errors"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BareMetalMachineProviderSpec holds data that the actuator needs to provision
// and manage a Machine.
// +k8s:openapi-gen=true
type BareMetalMachineProviderSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Image is the image to be provisioned.
	Image Image `json:"image,omitempty"`

	// UserData references the Secret that holds user data needed by the bare metal
	// operator. The Namespace is optional; it will default to the Machine's
	// namespace if not specified.
	UserData *corev1.SecretReference `json:"userData,omitempty"`

	// HostSelector specifies matching criteria for labels on BareMetalHosts.
	// This is used to limit the set of BareMetalHost objects considered for
	// claiming for a Machine.
	HostSelector HostSelector `json:"hostSelector,omitempty"`

	// BareMetalHostProfile specifies bare metal host configuration profile
	BareMetalHostProfile *BareMetalHostProfileReference `json:"bareMetalHostProfile,omitempty"`

	// AnsibleExtraPassthrough will be copied to bmh object without any modifications
	AnsibleExtraPassthrough *metal3v1alpha1.CustomJSON `json:"ansibleExtraPassthrough,omitempty"`

	// L2TemplateIfMappingOverride specified overrides for interface mapping
	// which will be used by IPAM.
	L2TemplateIfMappingOverride []string `json:"l2TemplateIfMappingOverride,omitempty"`

	// L2TemplateSelector will be passed to IpamHost while creation
	L2TemplateSelector *L2TemplateSelector `json:"l2TemplateSelector,omitempty"`

	kaasv1alpha1.MachineSpecMixin `json:",inline"`
}

type BareMetalHostProfileReference struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// L2TemplateSelector describes a criteria to select one of a cluster-related L2Template
type L2TemplateSelector struct {
	Name  string `json:"name,omitempty"`
	Label string `json:"label,omitempty"`
}

func (s *BareMetalMachineProviderSpec) GetMachineSpecMixin() *kaasv1alpha1.MachineSpecMixin {
	return &s.MachineSpecMixin
}

func (*BareMetalMachineProviderSpec) GetNewMachineStatus() runtime.Object {
	return &BareMetalMachineProviderStatus{}
}

// HostSelector specifies matching criteria for labels on BareMetalHosts.
// This is used to limit the set of BareMetalHost objects considered for
// claiming for a Machine.
type HostSelector struct {
	// Key/value pairs of labels that must exist on a chosen BareMetalHost
	MatchLabels map[string]string `json:"matchLabels,omitempty"`

	// Label match expressions that must be true on a chosen BareMetalHost
	MatchExpressions []HostSelectorRequirement `json:"matchExpressions,omitempty"`
}

type HostSelectorRequirement struct {
	Key      string             `json:"key"`
	Operator selection.Operator `json:"operator"`
	Values   []string           `json:"values"`
}

// Image holds the details of an image to use during provisioning.
type Image struct {
	// URL is a location of an image to deploy.
	URL string `json:"url,omitempty"`

	// Checksum is a md5sum value or a URL to retrieve one.
	Checksum string `json:"checksum,omitempty"`
}

// Validate returns an error if the object is not valid, otherwise nil. The
// string representation of the error is suitable for human consumption.
func (s *BareMetalMachineProviderSpec) Validate() error {
	missing := []string{}
	invalid := []string{}
	if s.Image.URL == "" {
		if imageURL, ok := os.LookupEnv("IRONIC_IMAGE_URL"); !ok {
			missing = append(missing, "Image.URL")
		} else {
			if !isValidURL(imageURL) {
				invalid = append(invalid, "Image.URL")
			}
		}
	}
	if s.Image.Checksum == "" {
		if imageChecksum, ok := os.LookupEnv("IRONIC_IMAGE_CHECKSUM"); !ok {
			missing = append(missing, "Image.Checksum")
		} else {
			if !isValidURL(imageChecksum) {
				invalid = append(invalid, "Image.Checksum")
			}
		}
	}
	if len(missing) > 0 {
		return errors.Errorf("Missing fields from ProviderSpec: %v", missing)
	}
	if len(invalid) > 0 {
		return errors.Errorf("Invalid URL format for fields from ProviderSpec: %v", invalid)
	}
	return nil
}

// IsValidURL returns an error if the object is not a valid URL consumable by BM provider,
// otherwise it returns true
func isValidURL(testURL string) bool {
	_, err := url.ParseRequestURI(testURL)
	if err != nil {
		return false
	}

	u, err := url.Parse(testURL)
	if err != nil || u.Scheme == "" || u.Host == "" || strings.HasPrefix(u.Host, ":") {
		return false
	}
	return true
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BareMetalMachineProviderSpecList contains a list of BareMetalMachineProviderSpec
type BareMetalMachineProviderSpecList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BareMetalMachineProviderSpec `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BareMetalMachineProviderSpec{}, &BareMetalMachineProviderSpecList{})
}

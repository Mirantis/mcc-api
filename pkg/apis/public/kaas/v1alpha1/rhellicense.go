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

// RHELLicense is the Schema for the Red Hat license API
// +k8s:openapi-gen=true
// +kubebuilder:resource
type RHELLicense struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RHELLicenseSpec `json:"spec,omitempty"`
}

type RHELLicenseSpec struct {
	Username string       `json:"username,omitempty"`
	Password *SecretValue `json:"password,omitempty"`

	OrgID         string       `json:"orgID,omitempty"`
	ActivationKey *SecretValue `json:"activationKey,omitempty"`

	PoolIDs []string `json:"poolIDs,omitempty"`
	RpmURL  string   `json:"rpmUrl,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RHELLicenseList contains a list of RHELLicense
type RHELLicenseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []RHELLicense `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RHELLicense{}, &RHELLicenseList{})
}

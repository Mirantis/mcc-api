/*
Copyright © 2020 Mirantis

Inspired by https://github.com/inwinstack/ipam/, https://github.com/inwinstack/blended/

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

package k8sutil

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kiConfig "github.com/Mirantis/mcc-api/pkg/apis/common/ipam/config"
	k8types "github.com/Mirantis/mcc-api/pkg/apis/util/ipam/k8sutil/types"
)

// GetRegionName -- returns region name of k8types.K8sObject or empty string if absent
func GetRegionName(m metav1.Object) string {
	if m == nil {
		return ""
	}
	return m.GetLabels()[kiConfig.RegionNameLabel]
}

// GetRegionNameFromContext -- returns region name from context
func GetRegionNameFromContext(ctx context.Context) (string, error) {
	regionName, ok := ctx.Value(kiConfig.RegionNameLabel).(string)
	if !ok {
		return "", fmt.Errorf("%w of region name, should be string: %#v", k8types.ErrorWrongFormat, ctx.Value(kiConfig.RegionNameLabel))
	} else if regionName == "" {
		return "", fmt.Errorf("wrong region name: empty string, %w", k8types.ErrorRegionUndefined)
	}
	return regionName, nil
}

// MachRegion -- if bool is true -- sring contains region name, else -- error message
func MachRegion(m metav1.Object) (string, bool) {
	regionName := GetRegionName(m)
	if regionName == "" {
		return "region name not given", false
	}
	if kiConfig.SingleRegionMode && regionName != kiConfig.ProviderConfig.Region {
		return "wrong region", false
	}
	return regionName, true
}

// GenerateRegionalName -- Return region-prefixed resource name
func GenerateRegionalName(name string) string {
	if !kiConfig.SingleRegionMode || kiConfig.ProviderConfig.Region == "" {
		return name
	}
	return kiConfig.ProviderConfig.Region + "-" + name
}

/*
Copyright © 2021 Mirantis

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
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kiConfig "github.com/Mirantis/mcc-api/pkg/apis/common/ipam/config"
	k8types "github.com/Mirantis/mcc-api/pkg/apis/util/ipam/k8sutil/types"
)

// EnsureUIDConsistency -- ensere that UIDlabel consistent to object UID
// rv indicates whether object was changed while ensuring
func EnsureUIDConsistency(m metav1.Object) (rv bool, err error) {
	if m == nil {
		return false, fmt.Errorf("%w: nil instead k8s object given", k8types.ErrorWrongParametr)
	}
	uid := string(m.GetUID())
	if uid == "" {
		return false, fmt.Errorf("%w: unable to get UID of k8s object %v", k8types.ErrorWrongObject, m)
	}
	labels := m.GetLabels()
	if labels == nil {
		labels = map[string]string{}
	}
	uidFromLabel := labels[kiConfig.UIDlabel]
	oldUIDlabel := fmt.Sprintf("%s-%s", kiConfig.UIDlabel, uid)
	_, oldUIDlabelFound := labels[oldUIDlabel]

	if uidFromLabel == "" {
		labels[kiConfig.UIDlabel] = uid
		labels[oldUIDlabel] = "1"
		rv = true
	}

	if !oldUIDlabelFound {
		labels[oldUIDlabel] = "1"
		rv = true
	}

	if uidFromLabel != uid && uidFromLabel != "" {
		labels[kiConfig.UIDlabel] = uid
		oldUIDlabel := fmt.Sprintf("%s-%s", kiConfig.UIDlabel, uidFromLabel)
		_, oldUIDlabelFound := labels[oldUIDlabel]
		if !oldUIDlabelFound {
			labels[oldUIDlabel] = "1"
		}
		rv = true
	}

	if rv {
		m.SetLabels(labels)
	}
	return rv, nil
}

/*
Copyright © 2021 Mirantis

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

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ListOptions(namespace string, matched map[string]string, exists []string) (*client.ListOptions, error) {
	optSelector := labels.NewSelector()

	for k := range matched {
		req, err := labels.NewRequirement(k, selection.Equals, []string{matched[k]})
		if err != nil {
			return nil, fmt.Errorf("error creating requirement for LabelSelector: %w", err)
		}
		optSelector = optSelector.Add(*req)
	}
	for i := range exists {
		req, err := labels.NewRequirement(exists[i], selection.Exists, nil)
		if err != nil {
			return nil, fmt.Errorf("error creating requirement for LabelSelector: %w", err)
		}
		optSelector = optSelector.Add(*req)
	}
	options := &client.ListOptions{
		Namespace:     namespace,
		LabelSelector: optSelector,
	}

	return options, nil
}

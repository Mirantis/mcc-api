package k8sutil

import (
	"fmt"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// +gocode:public-api=true
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

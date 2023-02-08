/*
Copyright 2022 The Kubernetes Authors.

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

// Code generated by client-gen@v0.20.2. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/Mirantis/mcc-api/v2/pkg/apis/kaas/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUnsupportedClusterses implements UnsupportedClustersInterface
type FakeUnsupportedClusterses struct {
	Fake *FakeKaasV1alpha1
}

var unsupportedclustersesResource = schema.GroupVersionResource{Group: "kaas.mirantis.com", Version: "v1alpha1", Resource: "unsupportedclusters"}

var unsupportedclustersesKind = schema.GroupVersionKind{Group: "kaas.mirantis.com", Version: "v1alpha1", Kind: "UnsupportedClusters"}

// Get takes name of the unsupportedClusters, and returns the corresponding unsupportedClusters object, and an error if there is any.
func (c *FakeUnsupportedClusterses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.UnsupportedClusters, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(unsupportedclustersesResource, name), &v1alpha1.UnsupportedClusters{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UnsupportedClusters), err
}

// List takes label and field selectors, and returns the list of UnsupportedClusterses that match those selectors.
func (c *FakeUnsupportedClusterses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.UnsupportedClustersList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(unsupportedclustersesResource, unsupportedclustersesKind, opts), &v1alpha1.UnsupportedClustersList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UnsupportedClustersList{ListMeta: obj.(*v1alpha1.UnsupportedClustersList).ListMeta}
	for _, item := range obj.(*v1alpha1.UnsupportedClustersList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested unsupportedClusterses.
func (c *FakeUnsupportedClusterses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(unsupportedclustersesResource, opts))
}

// Create takes the representation of a unsupportedClusters and creates it.  Returns the server's representation of the unsupportedClusters, and an error, if there is any.
func (c *FakeUnsupportedClusterses) Create(ctx context.Context, unsupportedClusters *v1alpha1.UnsupportedClusters, opts v1.CreateOptions) (result *v1alpha1.UnsupportedClusters, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(unsupportedclustersesResource, unsupportedClusters), &v1alpha1.UnsupportedClusters{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UnsupportedClusters), err
}

// Update takes the representation of a unsupportedClusters and updates it. Returns the server's representation of the unsupportedClusters, and an error, if there is any.
func (c *FakeUnsupportedClusterses) Update(ctx context.Context, unsupportedClusters *v1alpha1.UnsupportedClusters, opts v1.UpdateOptions) (result *v1alpha1.UnsupportedClusters, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(unsupportedclustersesResource, unsupportedClusters), &v1alpha1.UnsupportedClusters{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UnsupportedClusters), err
}

// Delete takes name of the unsupportedClusters and deletes it. Returns an error if one occurs.
func (c *FakeUnsupportedClusterses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(unsupportedclustersesResource, name), &v1alpha1.UnsupportedClusters{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUnsupportedClusterses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(unsupportedclustersesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.UnsupportedClustersList{})
	return err
}

// Patch applies the patch and returns the patched unsupportedClusters.
func (c *FakeUnsupportedClusterses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.UnsupportedClusters, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(unsupportedclustersesResource, name, pt, data, subresources...), &v1alpha1.UnsupportedClusters{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UnsupportedClusters), err
}

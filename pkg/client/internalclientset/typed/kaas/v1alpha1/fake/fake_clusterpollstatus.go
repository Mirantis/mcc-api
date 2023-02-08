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

// FakeClusterPollStatuses implements ClusterPollStatusInterface
type FakeClusterPollStatuses struct {
	Fake *FakeKaasV1alpha1
	ns   string
}

var clusterpollstatusesResource = schema.GroupVersionResource{Group: "kaas.mirantis.com", Version: "v1alpha1", Resource: "clusterpollstatuses"}

var clusterpollstatusesKind = schema.GroupVersionKind{Group: "kaas.mirantis.com", Version: "v1alpha1", Kind: "ClusterPollStatus"}

// Get takes name of the clusterPollStatus, and returns the corresponding clusterPollStatus object, and an error if there is any.
func (c *FakeClusterPollStatuses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterPollStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusterpollstatusesResource, c.ns, name), &v1alpha1.ClusterPollStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPollStatus), err
}

// List takes label and field selectors, and returns the list of ClusterPollStatuses that match those selectors.
func (c *FakeClusterPollStatuses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterPollStatusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusterpollstatusesResource, clusterpollstatusesKind, c.ns, opts), &v1alpha1.ClusterPollStatusList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterPollStatusList{ListMeta: obj.(*v1alpha1.ClusterPollStatusList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterPollStatusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterPollStatuses.
func (c *FakeClusterPollStatuses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusterpollstatusesResource, c.ns, opts))

}

// Create takes the representation of a clusterPollStatus and creates it.  Returns the server's representation of the clusterPollStatus, and an error, if there is any.
func (c *FakeClusterPollStatuses) Create(ctx context.Context, clusterPollStatus *v1alpha1.ClusterPollStatus, opts v1.CreateOptions) (result *v1alpha1.ClusterPollStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusterpollstatusesResource, c.ns, clusterPollStatus), &v1alpha1.ClusterPollStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPollStatus), err
}

// Update takes the representation of a clusterPollStatus and updates it. Returns the server's representation of the clusterPollStatus, and an error, if there is any.
func (c *FakeClusterPollStatuses) Update(ctx context.Context, clusterPollStatus *v1alpha1.ClusterPollStatus, opts v1.UpdateOptions) (result *v1alpha1.ClusterPollStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusterpollstatusesResource, c.ns, clusterPollStatus), &v1alpha1.ClusterPollStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPollStatus), err
}

// Delete takes name of the clusterPollStatus and deletes it. Returns an error if one occurs.
func (c *FakeClusterPollStatuses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clusterpollstatusesResource, c.ns, name), &v1alpha1.ClusterPollStatus{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterPollStatuses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusterpollstatusesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterPollStatusList{})
	return err
}

// Patch applies the patch and returns the patched clusterPollStatus.
func (c *FakeClusterPollStatuses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterPollStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusterpollstatusesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ClusterPollStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPollStatus), err
}

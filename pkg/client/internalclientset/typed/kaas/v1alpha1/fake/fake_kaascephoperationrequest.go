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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKaaSCephOperationRequests implements KaaSCephOperationRequestInterface
type FakeKaaSCephOperationRequests struct {
	Fake *FakeKaasV1alpha1
	ns   string
}

var kaascephoperationrequestsResource = schema.GroupVersionResource{Group: "kaas.mirantis.com", Version: "v1alpha1", Resource: "kaascephoperationrequests"}

var kaascephoperationrequestsKind = schema.GroupVersionKind{Group: "kaas.mirantis.com", Version: "v1alpha1", Kind: "KaaSCephOperationRequest"}

// Get takes name of the kaaSCephOperationRequest, and returns the corresponding kaaSCephOperationRequest object, and an error if there is any.
func (c *FakeKaaSCephOperationRequests) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KaaSCephOperationRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kaascephoperationrequestsResource, c.ns, name), &v1alpha1.KaaSCephOperationRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KaaSCephOperationRequest), err
}

// List takes label and field selectors, and returns the list of KaaSCephOperationRequests that match those selectors.
func (c *FakeKaaSCephOperationRequests) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KaaSCephOperationRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kaascephoperationrequestsResource, kaascephoperationrequestsKind, c.ns, opts), &v1alpha1.KaaSCephOperationRequestList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KaaSCephOperationRequestList{ListMeta: obj.(*v1alpha1.KaaSCephOperationRequestList).ListMeta}
	for _, item := range obj.(*v1alpha1.KaaSCephOperationRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kaaSCephOperationRequests.
func (c *FakeKaaSCephOperationRequests) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kaascephoperationrequestsResource, c.ns, opts))

}

// Create takes the representation of a kaaSCephOperationRequest and creates it.  Returns the server's representation of the kaaSCephOperationRequest, and an error, if there is any.
func (c *FakeKaaSCephOperationRequests) Create(ctx context.Context, kaaSCephOperationRequest *v1alpha1.KaaSCephOperationRequest, opts v1.CreateOptions) (result *v1alpha1.KaaSCephOperationRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kaascephoperationrequestsResource, c.ns, kaaSCephOperationRequest), &v1alpha1.KaaSCephOperationRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KaaSCephOperationRequest), err
}

// Update takes the representation of a kaaSCephOperationRequest and updates it. Returns the server's representation of the kaaSCephOperationRequest, and an error, if there is any.
func (c *FakeKaaSCephOperationRequests) Update(ctx context.Context, kaaSCephOperationRequest *v1alpha1.KaaSCephOperationRequest, opts v1.UpdateOptions) (result *v1alpha1.KaaSCephOperationRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kaascephoperationrequestsResource, c.ns, kaaSCephOperationRequest), &v1alpha1.KaaSCephOperationRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KaaSCephOperationRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKaaSCephOperationRequests) UpdateStatus(ctx context.Context, kaaSCephOperationRequest *v1alpha1.KaaSCephOperationRequest, opts v1.UpdateOptions) (*v1alpha1.KaaSCephOperationRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kaascephoperationrequestsResource, "status", c.ns, kaaSCephOperationRequest), &v1alpha1.KaaSCephOperationRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KaaSCephOperationRequest), err
}

// Delete takes name of the kaaSCephOperationRequest and deletes it. Returns an error if one occurs.
func (c *FakeKaaSCephOperationRequests) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kaascephoperationrequestsResource, c.ns, name), &v1alpha1.KaaSCephOperationRequest{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKaaSCephOperationRequests) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kaascephoperationrequestsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KaaSCephOperationRequestList{})
	return err
}

// Patch applies the patch and returns the patched kaaSCephOperationRequest.
func (c *FakeKaaSCephOperationRequests) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KaaSCephOperationRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kaascephoperationrequestsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KaaSCephOperationRequest{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KaaSCephOperationRequest), err
}

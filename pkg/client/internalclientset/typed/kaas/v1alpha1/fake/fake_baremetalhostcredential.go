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

// FakeBareMetalHostCredentials implements BareMetalHostCredentialInterface
type FakeBareMetalHostCredentials struct {
	Fake *FakeKaasV1alpha1
	ns   string
}

var baremetalhostcredentialsResource = schema.GroupVersionResource{Group: "kaas.mirantis.com", Version: "v1alpha1", Resource: "baremetalhostcredentials"}

var baremetalhostcredentialsKind = schema.GroupVersionKind{Group: "kaas.mirantis.com", Version: "v1alpha1", Kind: "BareMetalHostCredential"}

// Get takes name of the bareMetalHostCredential, and returns the corresponding bareMetalHostCredential object, and an error if there is any.
func (c *FakeBareMetalHostCredentials) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BareMetalHostCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(baremetalhostcredentialsResource, c.ns, name), &v1alpha1.BareMetalHostCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BareMetalHostCredential), err
}

// List takes label and field selectors, and returns the list of BareMetalHostCredentials that match those selectors.
func (c *FakeBareMetalHostCredentials) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BareMetalHostCredentialList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(baremetalhostcredentialsResource, baremetalhostcredentialsKind, c.ns, opts), &v1alpha1.BareMetalHostCredentialList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BareMetalHostCredentialList{ListMeta: obj.(*v1alpha1.BareMetalHostCredentialList).ListMeta}
	for _, item := range obj.(*v1alpha1.BareMetalHostCredentialList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bareMetalHostCredentials.
func (c *FakeBareMetalHostCredentials) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(baremetalhostcredentialsResource, c.ns, opts))

}

// Create takes the representation of a bareMetalHostCredential and creates it.  Returns the server's representation of the bareMetalHostCredential, and an error, if there is any.
func (c *FakeBareMetalHostCredentials) Create(ctx context.Context, bareMetalHostCredential *v1alpha1.BareMetalHostCredential, opts v1.CreateOptions) (result *v1alpha1.BareMetalHostCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(baremetalhostcredentialsResource, c.ns, bareMetalHostCredential), &v1alpha1.BareMetalHostCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BareMetalHostCredential), err
}

// Update takes the representation of a bareMetalHostCredential and updates it. Returns the server's representation of the bareMetalHostCredential, and an error, if there is any.
func (c *FakeBareMetalHostCredentials) Update(ctx context.Context, bareMetalHostCredential *v1alpha1.BareMetalHostCredential, opts v1.UpdateOptions) (result *v1alpha1.BareMetalHostCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(baremetalhostcredentialsResource, c.ns, bareMetalHostCredential), &v1alpha1.BareMetalHostCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BareMetalHostCredential), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBareMetalHostCredentials) UpdateStatus(ctx context.Context, bareMetalHostCredential *v1alpha1.BareMetalHostCredential, opts v1.UpdateOptions) (*v1alpha1.BareMetalHostCredential, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(baremetalhostcredentialsResource, "status", c.ns, bareMetalHostCredential), &v1alpha1.BareMetalHostCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BareMetalHostCredential), err
}

// Delete takes name of the bareMetalHostCredential and deletes it. Returns an error if one occurs.
func (c *FakeBareMetalHostCredentials) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(baremetalhostcredentialsResource, c.ns, name), &v1alpha1.BareMetalHostCredential{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBareMetalHostCredentials) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(baremetalhostcredentialsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BareMetalHostCredentialList{})
	return err
}

// Patch applies the patch and returns the patched bareMetalHostCredential.
func (c *FakeBareMetalHostCredentials) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BareMetalHostCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(baremetalhostcredentialsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BareMetalHostCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BareMetalHostCredential), err
}

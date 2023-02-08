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

// FakeOpenStackCredentials implements OpenStackCredentialInterface
type FakeOpenStackCredentials struct {
	Fake *FakeKaasV1alpha1
	ns   string
}

var openstackcredentialsResource = schema.GroupVersionResource{Group: "kaas.mirantis.com", Version: "v1alpha1", Resource: "openstackcredentials"}

var openstackcredentialsKind = schema.GroupVersionKind{Group: "kaas.mirantis.com", Version: "v1alpha1", Kind: "OpenStackCredential"}

// Get takes name of the openStackCredential, and returns the corresponding openStackCredential object, and an error if there is any.
func (c *FakeOpenStackCredentials) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OpenStackCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(openstackcredentialsResource, c.ns, name), &v1alpha1.OpenStackCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenStackCredential), err
}

// List takes label and field selectors, and returns the list of OpenStackCredentials that match those selectors.
func (c *FakeOpenStackCredentials) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OpenStackCredentialList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(openstackcredentialsResource, openstackcredentialsKind, c.ns, opts), &v1alpha1.OpenStackCredentialList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.OpenStackCredentialList{ListMeta: obj.(*v1alpha1.OpenStackCredentialList).ListMeta}
	for _, item := range obj.(*v1alpha1.OpenStackCredentialList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested openStackCredentials.
func (c *FakeOpenStackCredentials) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(openstackcredentialsResource, c.ns, opts))

}

// Create takes the representation of a openStackCredential and creates it.  Returns the server's representation of the openStackCredential, and an error, if there is any.
func (c *FakeOpenStackCredentials) Create(ctx context.Context, openStackCredential *v1alpha1.OpenStackCredential, opts v1.CreateOptions) (result *v1alpha1.OpenStackCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(openstackcredentialsResource, c.ns, openStackCredential), &v1alpha1.OpenStackCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenStackCredential), err
}

// Update takes the representation of a openStackCredential and updates it. Returns the server's representation of the openStackCredential, and an error, if there is any.
func (c *FakeOpenStackCredentials) Update(ctx context.Context, openStackCredential *v1alpha1.OpenStackCredential, opts v1.UpdateOptions) (result *v1alpha1.OpenStackCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(openstackcredentialsResource, c.ns, openStackCredential), &v1alpha1.OpenStackCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenStackCredential), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOpenStackCredentials) UpdateStatus(ctx context.Context, openStackCredential *v1alpha1.OpenStackCredential, opts v1.UpdateOptions) (*v1alpha1.OpenStackCredential, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(openstackcredentialsResource, "status", c.ns, openStackCredential), &v1alpha1.OpenStackCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenStackCredential), err
}

// Delete takes name of the openStackCredential and deletes it. Returns an error if one occurs.
func (c *FakeOpenStackCredentials) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(openstackcredentialsResource, c.ns, name), &v1alpha1.OpenStackCredential{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOpenStackCredentials) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(openstackcredentialsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.OpenStackCredentialList{})
	return err
}

// Patch applies the patch and returns the patched openStackCredential.
func (c *FakeOpenStackCredentials) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OpenStackCredential, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(openstackcredentialsResource, c.ns, name, pt, data, subresources...), &v1alpha1.OpenStackCredential{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.OpenStackCredential), err
}

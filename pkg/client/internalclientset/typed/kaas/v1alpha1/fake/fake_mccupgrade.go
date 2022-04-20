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

// FakeMCCUpgrades implements MCCUpgradeInterface
type FakeMCCUpgrades struct {
	Fake *FakeKaasV1alpha1
}

var mccupgradesResource = schema.GroupVersionResource{Group: "kaas.mirantis.com", Version: "v1alpha1", Resource: "mccupgrades"}

var mccupgradesKind = schema.GroupVersionKind{Group: "kaas.mirantis.com", Version: "v1alpha1", Kind: "MCCUpgrade"}

// Get takes name of the mCCUpgrade, and returns the corresponding mCCUpgrade object, and an error if there is any.
func (c *FakeMCCUpgrades) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MCCUpgrade, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(mccupgradesResource, name), &v1alpha1.MCCUpgrade{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MCCUpgrade), err
}

// List takes label and field selectors, and returns the list of MCCUpgrades that match those selectors.
func (c *FakeMCCUpgrades) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MCCUpgradeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(mccupgradesResource, mccupgradesKind, opts), &v1alpha1.MCCUpgradeList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MCCUpgradeList{ListMeta: obj.(*v1alpha1.MCCUpgradeList).ListMeta}
	for _, item := range obj.(*v1alpha1.MCCUpgradeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mCCUpgrades.
func (c *FakeMCCUpgrades) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(mccupgradesResource, opts))
}

// Create takes the representation of a mCCUpgrade and creates it.  Returns the server's representation of the mCCUpgrade, and an error, if there is any.
func (c *FakeMCCUpgrades) Create(ctx context.Context, mCCUpgrade *v1alpha1.MCCUpgrade, opts v1.CreateOptions) (result *v1alpha1.MCCUpgrade, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(mccupgradesResource, mCCUpgrade), &v1alpha1.MCCUpgrade{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MCCUpgrade), err
}

// Update takes the representation of a mCCUpgrade and updates it. Returns the server's representation of the mCCUpgrade, and an error, if there is any.
func (c *FakeMCCUpgrades) Update(ctx context.Context, mCCUpgrade *v1alpha1.MCCUpgrade, opts v1.UpdateOptions) (result *v1alpha1.MCCUpgrade, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(mccupgradesResource, mCCUpgrade), &v1alpha1.MCCUpgrade{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MCCUpgrade), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMCCUpgrades) UpdateStatus(ctx context.Context, mCCUpgrade *v1alpha1.MCCUpgrade, opts v1.UpdateOptions) (*v1alpha1.MCCUpgrade, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(mccupgradesResource, "status", mCCUpgrade), &v1alpha1.MCCUpgrade{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MCCUpgrade), err
}

// Delete takes name of the mCCUpgrade and deletes it. Returns an error if one occurs.
func (c *FakeMCCUpgrades) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(mccupgradesResource, name), &v1alpha1.MCCUpgrade{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMCCUpgrades) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(mccupgradesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MCCUpgradeList{})
	return err
}

// Patch applies the patch and returns the patched mCCUpgrade.
func (c *FakeMCCUpgrades) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MCCUpgrade, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(mccupgradesResource, name, pt, data, subresources...), &v1alpha1.MCCUpgrade{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MCCUpgrade), err
}

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

// FakeEquinixMetalResourceses implements EquinixMetalResourcesInterface
type FakeEquinixMetalResourceses struct {
	Fake *FakeKaasV1alpha1
	ns   string
}

var equinixmetalresourcesesResource = schema.GroupVersionResource{Group: "kaas.mirantis.com", Version: "v1alpha1", Resource: "equinixmetalresourceses"}

var equinixmetalresourcesesKind = schema.GroupVersionKind{Group: "kaas.mirantis.com", Version: "v1alpha1", Kind: "EquinixMetalResources"}

// Get takes name of the equinixMetalResources, and returns the corresponding equinixMetalResources object, and an error if there is any.
func (c *FakeEquinixMetalResourceses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EquinixMetalResources, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(equinixmetalresourcesesResource, c.ns, name), &v1alpha1.EquinixMetalResources{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EquinixMetalResources), err
}

// List takes label and field selectors, and returns the list of EquinixMetalResourceses that match those selectors.
func (c *FakeEquinixMetalResourceses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EquinixMetalResourcesList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(equinixmetalresourcesesResource, equinixmetalresourcesesKind, c.ns, opts), &v1alpha1.EquinixMetalResourcesList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EquinixMetalResourcesList{ListMeta: obj.(*v1alpha1.EquinixMetalResourcesList).ListMeta}
	for _, item := range obj.(*v1alpha1.EquinixMetalResourcesList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested equinixMetalResourceses.
func (c *FakeEquinixMetalResourceses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(equinixmetalresourcesesResource, c.ns, opts))

}

// Create takes the representation of a equinixMetalResources and creates it.  Returns the server's representation of the equinixMetalResources, and an error, if there is any.
func (c *FakeEquinixMetalResourceses) Create(ctx context.Context, equinixMetalResources *v1alpha1.EquinixMetalResources, opts v1.CreateOptions) (result *v1alpha1.EquinixMetalResources, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(equinixmetalresourcesesResource, c.ns, equinixMetalResources), &v1alpha1.EquinixMetalResources{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EquinixMetalResources), err
}

// Update takes the representation of a equinixMetalResources and updates it. Returns the server's representation of the equinixMetalResources, and an error, if there is any.
func (c *FakeEquinixMetalResourceses) Update(ctx context.Context, equinixMetalResources *v1alpha1.EquinixMetalResources, opts v1.UpdateOptions) (result *v1alpha1.EquinixMetalResources, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(equinixmetalresourcesesResource, c.ns, equinixMetalResources), &v1alpha1.EquinixMetalResources{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EquinixMetalResources), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEquinixMetalResourceses) UpdateStatus(ctx context.Context, equinixMetalResources *v1alpha1.EquinixMetalResources, opts v1.UpdateOptions) (*v1alpha1.EquinixMetalResources, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(equinixmetalresourcesesResource, "status", c.ns, equinixMetalResources), &v1alpha1.EquinixMetalResources{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EquinixMetalResources), err
}

// Delete takes name of the equinixMetalResources and deletes it. Returns an error if one occurs.
func (c *FakeEquinixMetalResourceses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(equinixmetalresourcesesResource, c.ns, name), &v1alpha1.EquinixMetalResources{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEquinixMetalResourceses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(equinixmetalresourcesesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EquinixMetalResourcesList{})
	return err
}

// Patch applies the patch and returns the patched equinixMetalResources.
func (c *FakeEquinixMetalResourceses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EquinixMetalResources, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(equinixmetalresourcesesResource, c.ns, name, pt, data, subresources...), &v1alpha1.EquinixMetalResources{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EquinixMetalResources), err
}
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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/Mirantis/mcc-api/pkg/apis/public/kaas/v1alpha1"
	scheme "github.com/Mirantis/mcc-api/pkg/client/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OpenStackResourcesesGetter has a method to return a OpenStackResourcesInterface.
// A group's client should implement this interface.
type OpenStackResourcesesGetter interface {
	OpenStackResourceses(namespace string) OpenStackResourcesInterface
}

// OpenStackResourcesInterface has methods to work with OpenStackResources resources.
type OpenStackResourcesInterface interface {
	Create(ctx context.Context, openStackResources *v1alpha1.OpenStackResources, opts v1.CreateOptions) (*v1alpha1.OpenStackResources, error)
	Update(ctx context.Context, openStackResources *v1alpha1.OpenStackResources, opts v1.UpdateOptions) (*v1alpha1.OpenStackResources, error)
	UpdateStatus(ctx context.Context, openStackResources *v1alpha1.OpenStackResources, opts v1.UpdateOptions) (*v1alpha1.OpenStackResources, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.OpenStackResources, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.OpenStackResourcesList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OpenStackResources, err error)
	OpenStackResourcesExpansion
}

// openStackResourceses implements OpenStackResourcesInterface
type openStackResourceses struct {
	client rest.Interface
	ns     string
}

// newOpenStackResourceses returns a OpenStackResourceses
func newOpenStackResourceses(c *KaasV1alpha1Client, namespace string) *openStackResourceses {
	return &openStackResourceses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the openStackResources, and returns the corresponding openStackResources object, and an error if there is any.
func (c *openStackResourceses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OpenStackResources, err error) {
	result = &v1alpha1.OpenStackResources{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("openstackresourceses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OpenStackResourceses that match those selectors.
func (c *openStackResourceses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OpenStackResourcesList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OpenStackResourcesList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("openstackresourceses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested openStackResourceses.
func (c *openStackResourceses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("openstackresourceses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a openStackResources and creates it.  Returns the server's representation of the openStackResources, and an error, if there is any.
func (c *openStackResourceses) Create(ctx context.Context, openStackResources *v1alpha1.OpenStackResources, opts v1.CreateOptions) (result *v1alpha1.OpenStackResources, err error) {
	result = &v1alpha1.OpenStackResources{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("openstackresourceses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(openStackResources).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a openStackResources and updates it. Returns the server's representation of the openStackResources, and an error, if there is any.
func (c *openStackResourceses) Update(ctx context.Context, openStackResources *v1alpha1.OpenStackResources, opts v1.UpdateOptions) (result *v1alpha1.OpenStackResources, err error) {
	result = &v1alpha1.OpenStackResources{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("openstackresourceses").
		Name(openStackResources.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(openStackResources).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *openStackResourceses) UpdateStatus(ctx context.Context, openStackResources *v1alpha1.OpenStackResources, opts v1.UpdateOptions) (result *v1alpha1.OpenStackResources, err error) {
	result = &v1alpha1.OpenStackResources{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("openstackresourceses").
		Name(openStackResources.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(openStackResources).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the openStackResources and deletes it. Returns an error if one occurs.
func (c *openStackResourceses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("openstackresourceses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *openStackResourceses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("openstackresourceses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched openStackResources.
func (c *openStackResourceses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OpenStackResources, err error) {
	result = &v1alpha1.OpenStackResources{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("openstackresourceses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

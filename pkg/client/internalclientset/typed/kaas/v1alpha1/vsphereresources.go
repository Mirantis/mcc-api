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

// VsphereResourcesesGetter has a method to return a VsphereResourcesInterface.
// A group's client should implement this interface.
type VsphereResourcesesGetter interface {
	VsphereResourceses(namespace string) VsphereResourcesInterface
}

// VsphereResourcesInterface has methods to work with VsphereResources resources.
type VsphereResourcesInterface interface {
	Create(ctx context.Context, vsphereResources *v1alpha1.VsphereResources, opts v1.CreateOptions) (*v1alpha1.VsphereResources, error)
	Update(ctx context.Context, vsphereResources *v1alpha1.VsphereResources, opts v1.UpdateOptions) (*v1alpha1.VsphereResources, error)
	UpdateStatus(ctx context.Context, vsphereResources *v1alpha1.VsphereResources, opts v1.UpdateOptions) (*v1alpha1.VsphereResources, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VsphereResources, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VsphereResourcesList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VsphereResources, err error)
	VsphereResourcesExpansion
}

// vsphereResourceses implements VsphereResourcesInterface
type vsphereResourceses struct {
	client rest.Interface
	ns     string
}

// newVsphereResourceses returns a VsphereResourceses
func newVsphereResourceses(c *KaasV1alpha1Client, namespace string) *vsphereResourceses {
	return &vsphereResourceses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vsphereResources, and returns the corresponding vsphereResources object, and an error if there is any.
func (c *vsphereResourceses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.VsphereResources, err error) {
	result = &v1alpha1.VsphereResources{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vsphereresourceses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VsphereResourceses that match those selectors.
func (c *vsphereResourceses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.VsphereResourcesList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VsphereResourcesList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vsphereresourceses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vsphereResourceses.
func (c *vsphereResourceses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vsphereresourceses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a vsphereResources and creates it.  Returns the server's representation of the vsphereResources, and an error, if there is any.
func (c *vsphereResourceses) Create(ctx context.Context, vsphereResources *v1alpha1.VsphereResources, opts v1.CreateOptions) (result *v1alpha1.VsphereResources, err error) {
	result = &v1alpha1.VsphereResources{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vsphereresourceses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vsphereResources).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a vsphereResources and updates it. Returns the server's representation of the vsphereResources, and an error, if there is any.
func (c *vsphereResourceses) Update(ctx context.Context, vsphereResources *v1alpha1.VsphereResources, opts v1.UpdateOptions) (result *v1alpha1.VsphereResources, err error) {
	result = &v1alpha1.VsphereResources{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vsphereresourceses").
		Name(vsphereResources.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vsphereResources).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *vsphereResourceses) UpdateStatus(ctx context.Context, vsphereResources *v1alpha1.VsphereResources, opts v1.UpdateOptions) (result *v1alpha1.VsphereResources, err error) {
	result = &v1alpha1.VsphereResources{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vsphereresourceses").
		Name(vsphereResources.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(vsphereResources).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the vsphereResources and deletes it. Returns an error if one occurs.
func (c *vsphereResourceses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vsphereresourceses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vsphereResourceses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vsphereresourceses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched vsphereResources.
func (c *vsphereResourceses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VsphereResources, err error) {
	result = &v1alpha1.VsphereResources{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vsphereresourceses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

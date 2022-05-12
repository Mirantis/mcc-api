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

// RHELLicensesGetter has a method to return a RHELLicenseInterface.
// A group's client should implement this interface.
type RHELLicensesGetter interface {
	RHELLicenses(namespace string) RHELLicenseInterface
}

// RHELLicenseInterface has methods to work with RHELLicense resources.
type RHELLicenseInterface interface {
	Create(ctx context.Context, rHELLicense *v1alpha1.RHELLicense, opts v1.CreateOptions) (*v1alpha1.RHELLicense, error)
	Update(ctx context.Context, rHELLicense *v1alpha1.RHELLicense, opts v1.UpdateOptions) (*v1alpha1.RHELLicense, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.RHELLicense, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.RHELLicenseList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RHELLicense, err error)
	RHELLicenseExpansion
}

// rHELLicenses implements RHELLicenseInterface
type rHELLicenses struct {
	client rest.Interface
	ns     string
}

// newRHELLicenses returns a RHELLicenses
func newRHELLicenses(c *KaasV1alpha1Client, namespace string) *rHELLicenses {
	return &rHELLicenses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the rHELLicense, and returns the corresponding rHELLicense object, and an error if there is any.
func (c *rHELLicenses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RHELLicense, err error) {
	result = &v1alpha1.RHELLicense{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rhellicenses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RHELLicenses that match those selectors.
func (c *rHELLicenses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RHELLicenseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RHELLicenseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rhellicenses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested rHELLicenses.
func (c *rHELLicenses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("rhellicenses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a rHELLicense and creates it.  Returns the server's representation of the rHELLicense, and an error, if there is any.
func (c *rHELLicenses) Create(ctx context.Context, rHELLicense *v1alpha1.RHELLicense, opts v1.CreateOptions) (result *v1alpha1.RHELLicense, err error) {
	result = &v1alpha1.RHELLicense{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("rhellicenses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(rHELLicense).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a rHELLicense and updates it. Returns the server's representation of the rHELLicense, and an error, if there is any.
func (c *rHELLicenses) Update(ctx context.Context, rHELLicense *v1alpha1.RHELLicense, opts v1.UpdateOptions) (result *v1alpha1.RHELLicense, err error) {
	result = &v1alpha1.RHELLicense{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rhellicenses").
		Name(rHELLicense.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(rHELLicense).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the rHELLicense and deletes it. Returns an error if one occurs.
func (c *rHELLicenses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rhellicenses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *rHELLicenses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rhellicenses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched rHELLicense.
func (c *rHELLicenses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RHELLicense, err error) {
	result = &v1alpha1.RHELLicense{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("rhellicenses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
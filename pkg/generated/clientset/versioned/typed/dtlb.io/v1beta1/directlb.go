/*
Copyright The Kubernetes Authors.

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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/huangjc7/directLB/pkg/apis/dtlb.io/v1beta1"
	scheme "github.com/huangjc7/directLB/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DirectLBsGetter has a method to return a DirectLBInterface.
// A group's client should implement this interface.
type DirectLBsGetter interface {
	DirectLBs(namespace string) DirectLBInterface
}

// DirectLBInterface has methods to work with DirectLB resources.
type DirectLBInterface interface {
	Create(ctx context.Context, directLB *v1beta1.DirectLB, opts v1.CreateOptions) (*v1beta1.DirectLB, error)
	Update(ctx context.Context, directLB *v1beta1.DirectLB, opts v1.UpdateOptions) (*v1beta1.DirectLB, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.DirectLB, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.DirectLBList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DirectLB, err error)
	DirectLBExpansion
}

// directLBs implements DirectLBInterface
type directLBs struct {
	client rest.Interface
	ns     string
}

// newDirectLBs returns a DirectLBs
func newDirectLBs(c *DtlbV1beta1Client, namespace string) *directLBs {
	return &directLBs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the directLB, and returns the corresponding directLB object, and an error if there is any.
func (c *directLBs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.DirectLB, err error) {
	result = &v1beta1.DirectLB{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("directlbs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DirectLBs that match those selectors.
func (c *directLBs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.DirectLBList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.DirectLBList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("directlbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested directLBs.
func (c *directLBs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("directlbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a directLB and creates it.  Returns the server's representation of the directLB, and an error, if there is any.
func (c *directLBs) Create(ctx context.Context, directLB *v1beta1.DirectLB, opts v1.CreateOptions) (result *v1beta1.DirectLB, err error) {
	result = &v1beta1.DirectLB{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("directlbs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(directLB).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a directLB and updates it. Returns the server's representation of the directLB, and an error, if there is any.
func (c *directLBs) Update(ctx context.Context, directLB *v1beta1.DirectLB, opts v1.UpdateOptions) (result *v1beta1.DirectLB, err error) {
	result = &v1beta1.DirectLB{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("directlbs").
		Name(directLB.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(directLB).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the directLB and deletes it. Returns an error if one occurs.
func (c *directLBs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("directlbs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *directLBs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("directlbs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched directLB.
func (c *directLBs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.DirectLB, err error) {
	result = &v1beta1.DirectLB{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("directlbs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

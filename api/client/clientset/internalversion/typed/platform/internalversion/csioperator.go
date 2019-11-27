/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
	platform "tkestack.io/tke/api/platform"
)

// CSIOperatorsGetter has a method to return a CSIOperatorInterface.
// A group's client should implement this interface.
type CSIOperatorsGetter interface {
	CSIOperators() CSIOperatorInterface
}

// CSIOperatorInterface has methods to work with CSIOperator resources.
type CSIOperatorInterface interface {
	Create(*platform.CSIOperator) (*platform.CSIOperator, error)
	Update(*platform.CSIOperator) (*platform.CSIOperator, error)
	UpdateStatus(*platform.CSIOperator) (*platform.CSIOperator, error)
	Delete(name string, options *v1.DeleteOptions) error
	Get(name string, options v1.GetOptions) (*platform.CSIOperator, error)
	List(opts v1.ListOptions) (*platform.CSIOperatorList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *platform.CSIOperator, err error)
	CSIOperatorExpansion
}

// cSIOperators implements CSIOperatorInterface
type cSIOperators struct {
	client rest.Interface
}

// newCSIOperators returns a CSIOperators
func newCSIOperators(c *PlatformClient) *cSIOperators {
	return &cSIOperators{
		client: c.RESTClient(),
	}
}

// Get takes name of the cSIOperator, and returns the corresponding cSIOperator object, and an error if there is any.
func (c *cSIOperators) Get(name string, options v1.GetOptions) (result *platform.CSIOperator, err error) {
	result = &platform.CSIOperator{}
	err = c.client.Get().
		Resource("csioperators").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CSIOperators that match those selectors.
func (c *cSIOperators) List(opts v1.ListOptions) (result *platform.CSIOperatorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &platform.CSIOperatorList{}
	err = c.client.Get().
		Resource("csioperators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cSIOperators.
func (c *cSIOperators) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("csioperators").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a cSIOperator and creates it.  Returns the server's representation of the cSIOperator, and an error, if there is any.
func (c *cSIOperators) Create(cSIOperator *platform.CSIOperator) (result *platform.CSIOperator, err error) {
	result = &platform.CSIOperator{}
	err = c.client.Post().
		Resource("csioperators").
		Body(cSIOperator).
		Do().
		Into(result)
	return
}

// Update takes the representation of a cSIOperator and updates it. Returns the server's representation of the cSIOperator, and an error, if there is any.
func (c *cSIOperators) Update(cSIOperator *platform.CSIOperator) (result *platform.CSIOperator, err error) {
	result = &platform.CSIOperator{}
	err = c.client.Put().
		Resource("csioperators").
		Name(cSIOperator.Name).
		Body(cSIOperator).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *cSIOperators) UpdateStatus(cSIOperator *platform.CSIOperator) (result *platform.CSIOperator, err error) {
	result = &platform.CSIOperator{}
	err = c.client.Put().
		Resource("csioperators").
		Name(cSIOperator.Name).
		SubResource("status").
		Body(cSIOperator).
		Do().
		Into(result)
	return
}

// Delete takes name of the cSIOperator and deletes it. Returns an error if one occurs.
func (c *cSIOperators) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("csioperators").
		Name(name).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched cSIOperator.
func (c *cSIOperators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *platform.CSIOperator, err error) {
	result = &platform.CSIOperator{}
	err = c.client.Patch(pt).
		Resource("csioperators").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
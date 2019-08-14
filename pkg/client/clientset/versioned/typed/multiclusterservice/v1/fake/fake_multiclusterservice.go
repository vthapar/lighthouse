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

package fake

import (
	multiclusterservicev1 "github.com/submariner-io/lighthouse/pkg/apis/multiclusterservice/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMultiClusterServices implements MultiClusterServiceInterface
type FakeMultiClusterServices struct {
	Fake *FakeLighthouseV1
	ns   string
}

var multiclusterservicesResource = schema.GroupVersionResource{Group: "lighthouse.submariner.io", Version: "v1", Resource: "multiclusterservices"}

var multiclusterservicesKind = schema.GroupVersionKind{Group: "lighthouse.submariner.io", Version: "v1", Kind: "MultiClusterService"}

// Get takes name of the multiClusterService, and returns the corresponding multiClusterService object, and an error if there is any.
func (c *FakeMultiClusterServices) Get(name string, options v1.GetOptions) (result *multiclusterservicev1.MultiClusterService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(multiclusterservicesResource, c.ns, name), &multiclusterservicev1.MultiClusterService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*multiclusterservicev1.MultiClusterService), err
}

// List takes label and field selectors, and returns the list of MultiClusterServices that match those selectors.
func (c *FakeMultiClusterServices) List(opts v1.ListOptions) (result *multiclusterservicev1.MultiClusterServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(multiclusterservicesResource, multiclusterservicesKind, c.ns, opts), &multiclusterservicev1.MultiClusterServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &multiclusterservicev1.MultiClusterServiceList{ListMeta: obj.(*multiclusterservicev1.MultiClusterServiceList).ListMeta}
	for _, item := range obj.(*multiclusterservicev1.MultiClusterServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested multiClusterServices.
func (c *FakeMultiClusterServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(multiclusterservicesResource, c.ns, opts))

}

// Create takes the representation of a multiClusterService and creates it.  Returns the server's representation of the multiClusterService, and an error, if there is any.
func (c *FakeMultiClusterServices) Create(multiClusterService *multiclusterservicev1.MultiClusterService) (result *multiclusterservicev1.MultiClusterService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(multiclusterservicesResource, c.ns, multiClusterService), &multiclusterservicev1.MultiClusterService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*multiclusterservicev1.MultiClusterService), err
}

// Update takes the representation of a multiClusterService and updates it. Returns the server's representation of the multiClusterService, and an error, if there is any.
func (c *FakeMultiClusterServices) Update(multiClusterService *multiclusterservicev1.MultiClusterService) (result *multiclusterservicev1.MultiClusterService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(multiclusterservicesResource, c.ns, multiClusterService), &multiclusterservicev1.MultiClusterService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*multiclusterservicev1.MultiClusterService), err
}

// Delete takes name of the multiClusterService and deletes it. Returns an error if one occurs.
func (c *FakeMultiClusterServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(multiclusterservicesResource, c.ns, name), &multiclusterservicev1.MultiClusterService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMultiClusterServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(multiclusterservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &multiclusterservicev1.MultiClusterServiceList{})
	return err
}

// Patch applies the patch and returns the patched multiClusterService.
func (c *FakeMultiClusterServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *multiclusterservicev1.MultiClusterService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(multiclusterservicesResource, c.ns, name, pt, data, subresources...), &multiclusterservicev1.MultiClusterService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*multiclusterservicev1.MultiClusterService), err
}
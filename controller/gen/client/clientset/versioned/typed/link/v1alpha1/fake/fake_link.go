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
	"context"

	v1alpha1 "github.com/linkerd/linkerd2/controller/gen/apis/link/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLinks implements LinkInterface
type FakeLinks struct {
	Fake *FakeLinkV1alpha1
	ns   string
}

var linksResource = schema.GroupVersionResource{Group: "multicluster.linkerd.io", Version: "v1alpha1", Resource: "links"}

var linksKind = schema.GroupVersionKind{Group: "multicluster.linkerd.io", Version: "v1alpha1", Kind: "Link"}

// Get takes name of the link, and returns the corresponding link object, and an error if there is any.
func (c *FakeLinks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Link, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(linksResource, c.ns, name), &v1alpha1.Link{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Link), err
}

// List takes label and field selectors, and returns the list of Links that match those selectors.
func (c *FakeLinks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LinkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(linksResource, linksKind, c.ns, opts), &v1alpha1.LinkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LinkList{ListMeta: obj.(*v1alpha1.LinkList).ListMeta}
	for _, item := range obj.(*v1alpha1.LinkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested links.
func (c *FakeLinks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(linksResource, c.ns, opts))

}

// Create takes the representation of a link and creates it.  Returns the server's representation of the link, and an error, if there is any.
func (c *FakeLinks) Create(ctx context.Context, link *v1alpha1.Link, opts v1.CreateOptions) (result *v1alpha1.Link, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(linksResource, c.ns, link), &v1alpha1.Link{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Link), err
}

// Update takes the representation of a link and updates it. Returns the server's representation of the link, and an error, if there is any.
func (c *FakeLinks) Update(ctx context.Context, link *v1alpha1.Link, opts v1.UpdateOptions) (result *v1alpha1.Link, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(linksResource, c.ns, link), &v1alpha1.Link{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Link), err
}

// Delete takes name of the link and deletes it. Returns an error if one occurs.
func (c *FakeLinks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(linksResource, c.ns, name), &v1alpha1.Link{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLinks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(linksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LinkList{})
	return err
}

// Patch applies the patch and returns the patched link.
func (c *FakeLinks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Link, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(linksResource, c.ns, name, pt, data, subresources...), &v1alpha1.Link{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Link), err
}

/*
   Copyright 2018 The Gitkube Authors

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
	v1alpha1 "github.com/thstarshine/gitkube/pkg/apis/gitkube.sh/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRemotes implements RemoteInterface
type FakeRemotes struct {
	Fake *FakeGitkubeV1alpha1
	ns   string
}

var remotesResource = schema.GroupVersionResource{Group: "gitkube.sh", Version: "v1alpha1", Resource: "remotes"}

var remotesKind = schema.GroupVersionKind{Group: "gitkube.sh", Version: "v1alpha1", Kind: "Remote"}

// Get takes name of the remote, and returns the corresponding remote object, and an error if there is any.
func (c *FakeRemotes) Get(name string, options v1.GetOptions) (result *v1alpha1.Remote, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(remotesResource, c.ns, name), &v1alpha1.Remote{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Remote), err
}

// List takes label and field selectors, and returns the list of Remotes that match those selectors.
func (c *FakeRemotes) List(opts v1.ListOptions) (result *v1alpha1.RemoteList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(remotesResource, remotesKind, c.ns, opts), &v1alpha1.RemoteList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RemoteList{ListMeta: obj.(*v1alpha1.RemoteList).ListMeta}
	for _, item := range obj.(*v1alpha1.RemoteList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested remotes.
func (c *FakeRemotes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(remotesResource, c.ns, opts))

}

// Create takes the representation of a remote and creates it.  Returns the server's representation of the remote, and an error, if there is any.
func (c *FakeRemotes) Create(remote *v1alpha1.Remote) (result *v1alpha1.Remote, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(remotesResource, c.ns, remote), &v1alpha1.Remote{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Remote), err
}

// Update takes the representation of a remote and updates it. Returns the server's representation of the remote, and an error, if there is any.
func (c *FakeRemotes) Update(remote *v1alpha1.Remote) (result *v1alpha1.Remote, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(remotesResource, c.ns, remote), &v1alpha1.Remote{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Remote), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRemotes) UpdateStatus(remote *v1alpha1.Remote) (*v1alpha1.Remote, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(remotesResource, "status", c.ns, remote), &v1alpha1.Remote{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Remote), err
}

// Delete takes name of the remote and deletes it. Returns an error if one occurs.
func (c *FakeRemotes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(remotesResource, c.ns, name), &v1alpha1.Remote{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRemotes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(remotesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.RemoteList{})
	return err
}

// Patch applies the patch and returns the patched remote.
func (c *FakeRemotes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Remote, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(remotesResource, c.ns, name, pt, data, subresources...), &v1alpha1.Remote{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Remote), err
}

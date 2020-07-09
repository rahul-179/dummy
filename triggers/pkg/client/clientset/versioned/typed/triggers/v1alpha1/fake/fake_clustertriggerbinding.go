/*
Copyright 2019 The Tekton Authors

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
	v1alpha1 "github.com/tektoncd/triggers/pkg/apis/triggers/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterTriggerBindings implements ClusterTriggerBindingInterface
type FakeClusterTriggerBindings struct {
	Fake *FakeTriggersV1alpha1
}

var clustertriggerbindingsResource = schema.GroupVersionResource{Group: "triggers.tekton.dev", Version: "v1alpha1", Resource: "clustertriggerbindings"}

var clustertriggerbindingsKind = schema.GroupVersionKind{Group: "triggers.tekton.dev", Version: "v1alpha1", Kind: "ClusterTriggerBinding"}

// Get takes name of the clusterTriggerBinding, and returns the corresponding clusterTriggerBinding object, and an error if there is any.
func (c *FakeClusterTriggerBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterTriggerBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clustertriggerbindingsResource, name), &v1alpha1.ClusterTriggerBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterTriggerBinding), err
}

// List takes label and field selectors, and returns the list of ClusterTriggerBindings that match those selectors.
func (c *FakeClusterTriggerBindings) List(opts v1.ListOptions) (result *v1alpha1.ClusterTriggerBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clustertriggerbindingsResource, clustertriggerbindingsKind, opts), &v1alpha1.ClusterTriggerBindingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterTriggerBindingList{ListMeta: obj.(*v1alpha1.ClusterTriggerBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterTriggerBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterTriggerBindings.
func (c *FakeClusterTriggerBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clustertriggerbindingsResource, opts))
}

// Create takes the representation of a clusterTriggerBinding and creates it.  Returns the server's representation of the clusterTriggerBinding, and an error, if there is any.
func (c *FakeClusterTriggerBindings) Create(clusterTriggerBinding *v1alpha1.ClusterTriggerBinding) (result *v1alpha1.ClusterTriggerBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clustertriggerbindingsResource, clusterTriggerBinding), &v1alpha1.ClusterTriggerBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterTriggerBinding), err
}

// Update takes the representation of a clusterTriggerBinding and updates it. Returns the server's representation of the clusterTriggerBinding, and an error, if there is any.
func (c *FakeClusterTriggerBindings) Update(clusterTriggerBinding *v1alpha1.ClusterTriggerBinding) (result *v1alpha1.ClusterTriggerBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clustertriggerbindingsResource, clusterTriggerBinding), &v1alpha1.ClusterTriggerBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterTriggerBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterTriggerBindings) UpdateStatus(clusterTriggerBinding *v1alpha1.ClusterTriggerBinding) (*v1alpha1.ClusterTriggerBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clustertriggerbindingsResource, "status", clusterTriggerBinding), &v1alpha1.ClusterTriggerBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterTriggerBinding), err
}

// Delete takes name of the clusterTriggerBinding and deletes it. Returns an error if one occurs.
func (c *FakeClusterTriggerBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clustertriggerbindingsResource, name), &v1alpha1.ClusterTriggerBinding{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterTriggerBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clustertriggerbindingsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterTriggerBindingList{})
	return err
}

// Patch applies the patch and returns the patched clusterTriggerBinding.
func (c *FakeClusterTriggerBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterTriggerBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clustertriggerbindingsResource, name, pt, data, subresources...), &v1alpha1.ClusterTriggerBinding{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterTriggerBinding), err
}

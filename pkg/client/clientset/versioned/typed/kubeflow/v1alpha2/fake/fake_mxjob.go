// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha2 "github.com/kubeflow/mxnet-operator.v2/pkg/apis/mxnet/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMXJobs implements MXJobInterface
type FakeMXJobs struct {
	Fake *FakeKubeflowV1alpha2
	ns   string
}

var mxjobsResource = schema.GroupVersionResource{Group: "kubeflow.org", Version: "v1alpha2", Resource: "mxjobs"}

var mxjobsKind = schema.GroupVersionKind{Group: "kubeflow.org", Version: "v1alpha2", Kind: "MXJob"}

// Get takes name of the mXJob, and returns the corresponding mXJob object, and an error if there is any.
func (c *FakeMXJobs) Get(name string, options v1.GetOptions) (result *v1alpha2.MXJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mxjobsResource, c.ns, name), &v1alpha2.MXJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.MXJob), err
}

// List takes label and field selectors, and returns the list of MXJobs that match those selectors.
func (c *FakeMXJobs) List(opts v1.ListOptions) (result *v1alpha2.MXJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mxjobsResource, mxjobsKind, c.ns, opts), &v1alpha2.MXJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.MXJobList{}
	for _, item := range obj.(*v1alpha2.MXJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mXJobs.
func (c *FakeMXJobs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mxjobsResource, c.ns, opts))

}

// Create takes the representation of a mXJob and creates it.  Returns the server's representation of the mXJob, and an error, if there is any.
func (c *FakeMXJobs) Create(mXJob *v1alpha2.MXJob) (result *v1alpha2.MXJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mxjobsResource, c.ns, mXJob), &v1alpha2.MXJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.MXJob), err
}

// Update takes the representation of a mXJob and updates it. Returns the server's representation of the mXJob, and an error, if there is any.
func (c *FakeMXJobs) Update(mXJob *v1alpha2.MXJob) (result *v1alpha2.MXJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mxjobsResource, c.ns, mXJob), &v1alpha2.MXJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.MXJob), err
}

// Delete takes name of the mXJob and deletes it. Returns an error if one occurs.
func (c *FakeMXJobs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mxjobsResource, c.ns, name), &v1alpha2.MXJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMXJobs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mxjobsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.MXJobList{})
	return err
}

// Patch applies the patch and returns the patched mXJob.
func (c *FakeMXJobs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.MXJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mxjobsResource, c.ns, name, data, subresources...), &v1alpha2.MXJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.MXJob), err
}

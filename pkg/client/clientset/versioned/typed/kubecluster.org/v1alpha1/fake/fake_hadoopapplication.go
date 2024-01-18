/*
Copyright 2023.

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

	v1alpha1 "github.com/chriskery/hadoop-operator/pkg/apis/kubecluster.org/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHadoopApplications implements HadoopApplicationInterface
type FakeHadoopApplications struct {
	Fake *FakeKubeclusterV1alpha1
	ns   string
}

var hadoopapplicationsResource = v1alpha1.SchemeGroupVersion.WithResource("hadoopapplications")

var hadoopapplicationsKind = v1alpha1.SchemeGroupVersion.WithKind("HadoopApplication")

// Get takes name of the hadoopApplication, and returns the corresponding hadoopApplication object, and an error if there is any.
func (c *FakeHadoopApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HadoopApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hadoopapplicationsResource, c.ns, name), &v1alpha1.HadoopApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HadoopApplication), err
}

// List takes label and field selectors, and returns the list of HadoopApplications that match those selectors.
func (c *FakeHadoopApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HadoopApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hadoopapplicationsResource, hadoopapplicationsKind, c.ns, opts), &v1alpha1.HadoopApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HadoopApplicationList{ListMeta: obj.(*v1alpha1.HadoopApplicationList).ListMeta}
	for _, item := range obj.(*v1alpha1.HadoopApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hadoopApplications.
func (c *FakeHadoopApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hadoopapplicationsResource, c.ns, opts))

}

// Create takes the representation of a hadoopApplication and creates it.  Returns the server's representation of the hadoopApplication, and an error, if there is any.
func (c *FakeHadoopApplications) Create(ctx context.Context, hadoopApplication *v1alpha1.HadoopApplication, opts v1.CreateOptions) (result *v1alpha1.HadoopApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hadoopapplicationsResource, c.ns, hadoopApplication), &v1alpha1.HadoopApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HadoopApplication), err
}

// Update takes the representation of a hadoopApplication and updates it. Returns the server's representation of the hadoopApplication, and an error, if there is any.
func (c *FakeHadoopApplications) Update(ctx context.Context, hadoopApplication *v1alpha1.HadoopApplication, opts v1.UpdateOptions) (result *v1alpha1.HadoopApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hadoopapplicationsResource, c.ns, hadoopApplication), &v1alpha1.HadoopApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HadoopApplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHadoopApplications) UpdateStatus(ctx context.Context, hadoopApplication *v1alpha1.HadoopApplication, opts v1.UpdateOptions) (*v1alpha1.HadoopApplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hadoopapplicationsResource, "status", c.ns, hadoopApplication), &v1alpha1.HadoopApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HadoopApplication), err
}

// Delete takes name of the hadoopApplication and deletes it. Returns an error if one occurs.
func (c *FakeHadoopApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(hadoopapplicationsResource, c.ns, name, opts), &v1alpha1.HadoopApplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHadoopApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hadoopapplicationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HadoopApplicationList{})
	return err
}

// Patch applies the patch and returns the patched hadoopApplication.
func (c *FakeHadoopApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HadoopApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hadoopapplicationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.HadoopApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HadoopApplication), err
}

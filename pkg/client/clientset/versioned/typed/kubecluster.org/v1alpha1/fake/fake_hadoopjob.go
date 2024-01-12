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

	v1alpha1 "github.com/chriskery/hadoop-cluster-operator/pkg/apis/kubecluster.org/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHadoopJobs implements HadoopJobInterface
type FakeHadoopJobs struct {
	Fake *FakeKubeclusterV1alpha1
	ns   string
}

var hadoopjobsResource = v1alpha1.SchemeGroupVersion.WithResource("hadoopjobs")

var hadoopjobsKind = v1alpha1.SchemeGroupVersion.WithKind("HadoopJob")

// Get takes name of the hadoopJob, and returns the corresponding hadoopJob object, and an error if there is any.
func (c *FakeHadoopJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HadoopJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hadoopjobsResource, c.ns, name), &v1alpha1.HadoopJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HadoopJob), err
}

// List takes label and field selectors, and returns the list of HadoopJobs that match those selectors.
func (c *FakeHadoopJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HadoopJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hadoopjobsResource, hadoopjobsKind, c.ns, opts), &v1alpha1.HadoopJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HadoopJobList{ListMeta: obj.(*v1alpha1.HadoopJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.HadoopJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hadoopJobs.
func (c *FakeHadoopJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hadoopjobsResource, c.ns, opts))

}

// Create takes the representation of a hadoopJob and creates it.  Returns the server's representation of the hadoopJob, and an error, if there is any.
func (c *FakeHadoopJobs) Create(ctx context.Context, hadoopJob *v1alpha1.HadoopJob, opts v1.CreateOptions) (result *v1alpha1.HadoopJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hadoopjobsResource, c.ns, hadoopJob), &v1alpha1.HadoopJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HadoopJob), err
}

// Update takes the representation of a hadoopJob and updates it. Returns the server's representation of the hadoopJob, and an error, if there is any.
func (c *FakeHadoopJobs) Update(ctx context.Context, hadoopJob *v1alpha1.HadoopJob, opts v1.UpdateOptions) (result *v1alpha1.HadoopJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hadoopjobsResource, c.ns, hadoopJob), &v1alpha1.HadoopJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HadoopJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHadoopJobs) UpdateStatus(ctx context.Context, hadoopJob *v1alpha1.HadoopJob, opts v1.UpdateOptions) (*v1alpha1.HadoopJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hadoopjobsResource, "status", c.ns, hadoopJob), &v1alpha1.HadoopJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HadoopJob), err
}

// Delete takes name of the hadoopJob and deletes it. Returns an error if one occurs.
func (c *FakeHadoopJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(hadoopjobsResource, c.ns, name, opts), &v1alpha1.HadoopJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHadoopJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hadoopjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HadoopJobList{})
	return err
}

// Patch applies the patch and returns the patched hadoopJob.
func (c *FakeHadoopJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HadoopJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hadoopjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.HadoopJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HadoopJob), err
}

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

// FakeHadoopClusters implements HadoopClusterInterface
type FakeHadoopClusters struct {
	Fake *FakeKubeclusterV1alpha1
	ns   string
}

var hadoopclustersResource = v1alpha1.SchemeGroupVersion.WithResource("hadoopclusters")

var hadoopclustersKind = v1alpha1.SchemeGroupVersion.WithKind("HadoopCluster")

// Get takes name of the hadoopCluster, and returns the corresponding hadoopCluster object, and an error if there is any.
func (c *FakeHadoopClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HadoopCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hadoopclustersResource, c.ns, name), &v1alpha1.HadoopCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HadoopCluster), err
}

// List takes label and field selectors, and returns the list of HadoopClusters that match those selectors.
func (c *FakeHadoopClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HadoopClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hadoopclustersResource, hadoopclustersKind, c.ns, opts), &v1alpha1.HadoopClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HadoopClusterList{ListMeta: obj.(*v1alpha1.HadoopClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.HadoopClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hadoopClusters.
func (c *FakeHadoopClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hadoopclustersResource, c.ns, opts))

}

// Create takes the representation of a hadoopCluster and creates it.  Returns the server's representation of the hadoopCluster, and an error, if there is any.
func (c *FakeHadoopClusters) Create(ctx context.Context, hadoopCluster *v1alpha1.HadoopCluster, opts v1.CreateOptions) (result *v1alpha1.HadoopCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hadoopclustersResource, c.ns, hadoopCluster), &v1alpha1.HadoopCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HadoopCluster), err
}

// Update takes the representation of a hadoopCluster and updates it. Returns the server's representation of the hadoopCluster, and an error, if there is any.
func (c *FakeHadoopClusters) Update(ctx context.Context, hadoopCluster *v1alpha1.HadoopCluster, opts v1.UpdateOptions) (result *v1alpha1.HadoopCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hadoopclustersResource, c.ns, hadoopCluster), &v1alpha1.HadoopCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HadoopCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHadoopClusters) UpdateStatus(ctx context.Context, hadoopCluster *v1alpha1.HadoopCluster, opts v1.UpdateOptions) (*v1alpha1.HadoopCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hadoopclustersResource, "status", c.ns, hadoopCluster), &v1alpha1.HadoopCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HadoopCluster), err
}

// Delete takes name of the hadoopCluster and deletes it. Returns an error if one occurs.
func (c *FakeHadoopClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(hadoopclustersResource, c.ns, name, opts), &v1alpha1.HadoopCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHadoopClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hadoopclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HadoopClusterList{})
	return err
}

// Patch applies the patch and returns the patched hadoopCluster.
func (c *FakeHadoopClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HadoopCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hadoopclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.HadoopCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HadoopCluster), err
}

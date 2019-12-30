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
	txzdatahubv1 "datahub.txzing.com/mysql-gr-operator/pkg/apis/txz_datahub/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMysqlGROperators implements MysqlGROperatorInterface
type FakeMysqlGROperators struct {
	Fake *FakeDatahubV1
	ns   string
}

var mysqlgroperatorsResource = schema.GroupVersionResource{Group: "datahub.txzing.com", Version: "v1", Resource: "mysqlgroperators"}

var mysqlgroperatorsKind = schema.GroupVersionKind{Group: "datahub.txzing.com", Version: "v1", Kind: "MysqlGROperator"}

// Get takes name of the mysqlGROperator, and returns the corresponding mysqlGROperator object, and an error if there is any.
func (c *FakeMysqlGROperators) Get(name string, options v1.GetOptions) (result *txzdatahubv1.MysqlGROperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mysqlgroperatorsResource, c.ns, name), &txzdatahubv1.MysqlGROperator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*txzdatahubv1.MysqlGROperator), err
}

// List takes label and field selectors, and returns the list of MysqlGROperators that match those selectors.
func (c *FakeMysqlGROperators) List(opts v1.ListOptions) (result *txzdatahubv1.MysqlGROperatorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mysqlgroperatorsResource, mysqlgroperatorsKind, c.ns, opts), &txzdatahubv1.MysqlGROperatorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &txzdatahubv1.MysqlGROperatorList{ListMeta: obj.(*txzdatahubv1.MysqlGROperatorList).ListMeta}
	for _, item := range obj.(*txzdatahubv1.MysqlGROperatorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mysqlGROperators.
func (c *FakeMysqlGROperators) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mysqlgroperatorsResource, c.ns, opts))

}

// Create takes the representation of a mysqlGROperator and creates it.  Returns the server's representation of the mysqlGROperator, and an error, if there is any.
func (c *FakeMysqlGROperators) Create(mysqlGROperator *txzdatahubv1.MysqlGROperator) (result *txzdatahubv1.MysqlGROperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mysqlgroperatorsResource, c.ns, mysqlGROperator), &txzdatahubv1.MysqlGROperator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*txzdatahubv1.MysqlGROperator), err
}

// Update takes the representation of a mysqlGROperator and updates it. Returns the server's representation of the mysqlGROperator, and an error, if there is any.
func (c *FakeMysqlGROperators) Update(mysqlGROperator *txzdatahubv1.MysqlGROperator) (result *txzdatahubv1.MysqlGROperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mysqlgroperatorsResource, c.ns, mysqlGROperator), &txzdatahubv1.MysqlGROperator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*txzdatahubv1.MysqlGROperator), err
}

// Delete takes name of the mysqlGROperator and deletes it. Returns an error if one occurs.
func (c *FakeMysqlGROperators) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mysqlgroperatorsResource, c.ns, name), &txzdatahubv1.MysqlGROperator{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMysqlGROperators) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mysqlgroperatorsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &txzdatahubv1.MysqlGROperatorList{})
	return err
}

// Patch applies the patch and returns the patched mysqlGROperator.
func (c *FakeMysqlGROperators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *txzdatahubv1.MysqlGROperator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mysqlgroperatorsResource, c.ns, name, pt, data, subresources...), &txzdatahubv1.MysqlGROperator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*txzdatahubv1.MysqlGROperator), err
}

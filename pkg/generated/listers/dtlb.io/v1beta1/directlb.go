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
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/huangjc7/directLB/pkg/apis/dtlb.io/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DirectLBLister helps list DirectLBs.
// All objects returned here must be treated as read-only.
type DirectLBLister interface {
	// List lists all DirectLBs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.DirectLB, err error)
	// DirectLBs returns an object that can list and get DirectLBs.
	DirectLBs(namespace string) DirectLBNamespaceLister
	DirectLBListerExpansion
}

// directLBLister implements the DirectLBLister interface.
type directLBLister struct {
	indexer cache.Indexer
}

// NewDirectLBLister returns a new DirectLBLister.
func NewDirectLBLister(indexer cache.Indexer) DirectLBLister {
	return &directLBLister{indexer: indexer}
}

// List lists all DirectLBs in the indexer.
func (s *directLBLister) List(selector labels.Selector) (ret []*v1beta1.DirectLB, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.DirectLB))
	})
	return ret, err
}

// DirectLBs returns an object that can list and get DirectLBs.
func (s *directLBLister) DirectLBs(namespace string) DirectLBNamespaceLister {
	return directLBNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DirectLBNamespaceLister helps list and get DirectLBs.
// All objects returned here must be treated as read-only.
type DirectLBNamespaceLister interface {
	// List lists all DirectLBs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.DirectLB, err error)
	// Get retrieves the DirectLB from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.DirectLB, error)
	DirectLBNamespaceListerExpansion
}

// directLBNamespaceLister implements the DirectLBNamespaceLister
// interface.
type directLBNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DirectLBs in the indexer for a given namespace.
func (s directLBNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.DirectLB, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.DirectLB))
	})
	return ret, err
}

// Get retrieves the DirectLB from the indexer for a given namespace and name.
func (s directLBNamespaceLister) Get(name string) (*v1beta1.DirectLB, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("directlb"), name)
	}
	return obj.(*v1beta1.DirectLB), nil
}

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
// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	dtlbiov1beta1 "github.com/huangjc7/directLB/pkg/apis/dtlb.io/v1beta1"
	versioned "github.com/huangjc7/directLB/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/huangjc7/directLB/pkg/generated/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/huangjc7/directLB/pkg/generated/listers/dtlb.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// DirectLBInformer provides access to a shared informer and lister for
// DirectLBs.
type DirectLBInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.DirectLBLister
}

type directLBInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewDirectLBInformer constructs a new informer for DirectLB type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewDirectLBInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredDirectLBInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredDirectLBInformer constructs a new informer for DirectLB type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredDirectLBInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DtlbV1beta1().DirectLBs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DtlbV1beta1().DirectLBs(namespace).Watch(context.TODO(), options)
			},
		},
		&dtlbiov1beta1.DirectLB{},
		resyncPeriod,
		indexers,
	)
}

func (f *directLBInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredDirectLBInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *directLBInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&dtlbiov1beta1.DirectLB{}, f.defaultInformer)
}

func (f *directLBInformer) Lister() v1beta1.DirectLBLister {
	return v1beta1.NewDirectLBLister(f.Informer().GetIndexer())
}

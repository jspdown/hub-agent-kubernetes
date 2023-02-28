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

package v1alpha1

import (
	"context"
	time "time"

	hubv1alpha1 "github.com/traefik/hub-agent-kubernetes/pkg/crd/api/hub/v1alpha1"
	versioned "github.com/traefik/hub-agent-kubernetes/pkg/crd/generated/client/hub/clientset/versioned"
	internalinterfaces "github.com/traefik/hub-agent-kubernetes/pkg/crd/generated/client/hub/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/traefik/hub-agent-kubernetes/pkg/crd/generated/client/hub/listers/hub/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// APIAccessInformer provides access to a shared informer and lister for
// APIAccesses.
type APIAccessInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.APIAccessLister
}

type aPIAccessInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAPIAccessInformer constructs a new informer for APIAccess type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAPIAccessInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAPIAccessInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAPIAccessInformer constructs a new informer for APIAccess type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAPIAccessInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HubV1alpha1().APIAccesses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HubV1alpha1().APIAccesses(namespace).Watch(context.TODO(), options)
			},
		},
		&hubv1alpha1.APIAccess{},
		resyncPeriod,
		indexers,
	)
}

func (f *aPIAccessInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAPIAccessInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *aPIAccessInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hubv1alpha1.APIAccess{}, f.defaultInformer)
}

func (f *aPIAccessInformer) Lister() v1alpha1.APIAccessLister {
	return v1alpha1.NewAPIAccessLister(f.Informer().GetIndexer())
}

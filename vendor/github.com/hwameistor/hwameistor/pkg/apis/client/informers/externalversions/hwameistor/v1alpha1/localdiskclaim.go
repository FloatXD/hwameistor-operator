// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	versioned "github.com/hwameistor/hwameistor/pkg/apis/client/clientset/versioned"
	internalinterfaces "github.com/hwameistor/hwameistor/pkg/apis/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/client/listers/hwameistor/v1alpha1"
	hwameistorv1alpha1 "github.com/hwameistor/hwameistor/pkg/apis/hwameistor/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LocalDiskClaimInformer provides access to a shared informer and lister for
// LocalDiskClaims.
type LocalDiskClaimInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.LocalDiskClaimLister
}

type localDiskClaimInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewLocalDiskClaimInformer constructs a new informer for LocalDiskClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLocalDiskClaimInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLocalDiskClaimInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredLocalDiskClaimInformer constructs a new informer for LocalDiskClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLocalDiskClaimInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HwameistorV1alpha1().LocalDiskClaims().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.HwameistorV1alpha1().LocalDiskClaims().Watch(context.TODO(), options)
			},
		},
		&hwameistorv1alpha1.LocalDiskClaim{},
		resyncPeriod,
		indexers,
	)
}

func (f *localDiskClaimInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLocalDiskClaimInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *localDiskClaimInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&hwameistorv1alpha1.LocalDiskClaim{}, f.defaultInformer)
}

func (f *localDiskClaimInformer) Lister() v1alpha1.LocalDiskClaimLister {
	return v1alpha1.NewLocalDiskClaimLister(f.Informer().GetIndexer())
}

package platform

import (
	"context"

	"github.com/rs/zerolog/log"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterResource interface {
	GetObjectMeta() metav1.ObjectMeta
}

type PlatformResource[T ClusterResource] interface {
	Resource() (T, error)
}

type ResourceHandler[T ClusterResource] interface {
	List() ([]T, error)
	Create(ctx context.Context, new ClusterResource) error
	Update(ctx context.Context, old, new ClusterResource) error
	Delete(ctx context.Context, old ClusterResource) error
}

type ResourceLister[T any] func(ctx context.Context) ([]T, error)

type ResourceSyncer[C ClusterResource, P PlatformResource[C]] struct {
	listPlatformResources ResourceLister[P]
	cluster               ResourceHandler[C]
}

func NewResourceSyncer[C ClusterResource, P PlatformResource[C]](
	listPlatformResources ResourceLister[P],
	cluster ResourceHandler[C]) *ResourceSyncer[C, P] {
	return &ResourceSyncer[C, P]{
		listPlatformResources: listPlatformResources,
		cluster:               cluster,
	}
}

func (r *ResourceSyncer[C, P]) Sync(ctx context.Context) {
	platformResources, err := r.listPlatformResources(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Unable to fetch KIND")
		return
	}

	clusterResources, err := r.cluster.List()
	if err != nil {
		log.Error().Err(err).Msg("Unable to fetch KIND")
		return
	}

	clusterResourceByName := make(map[string]C)
	for _, resource := range clusterResources {
		clusterResourceByName[resourceKey(resource)] = resource
	}

	for _, platformResource := range platformResources {
		newClusterResource, err := platformResource.Resource()
		if err != nil {
			log.Error().Err(err).Msg("Unable to build cluster resource")
			continue
		}

		key := resourceKey(newClusterResource)
		oldClusterResource, found := clusterResourceByName[key]

		// Resources that will remain in the map will be deleted.
		delete(clusterResourceByName, key)

		if !found {
			if err = r.cluster.Create(ctx, newClusterResource); err != nil {
				log.Error().Err(err).
					Str("name", key).
					Msg("Unable to create KIND")
			}
			continue
		}

		if err = r.cluster.Update(ctx, oldClusterResource, newClusterResource); err != nil {
			log.Error().Err(err).
				Str("name", key).
				Msg("Unable to update KIND")
		}
	}

	for _, clusterResource := range clusterResourceByName {
		if err = r.cluster.Delete(ctx, clusterResource); err != nil {
			log.Error().Err(err).Msg("Unable to delete KIND")

			continue
		}
	}
}

func resourceKey(resource ClusterResource) string {
	meta := resource.GetObjectMeta()

	if meta.Namespace != "" {
		return meta.Name + "@" + meta.Namespace
	}
	return meta.Name
}

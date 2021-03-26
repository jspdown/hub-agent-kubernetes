package ingclass

import (
	"errors"
	"fmt"
	"sync"

	"github.com/rs/zerolog/log"
	neov1alpha1 "github.com/traefik/neo-agent/pkg/crd/api/neo/v1alpha1"
	netv1 "k8s.io/api/networking/v1"
	netv1beta1 "k8s.io/api/networking/v1beta1"
	ktypes "k8s.io/apimachinery/pkg/types"
)

// ingressClass is an internal representation of either a netv1.IngressClass,
// netv1beta1.IngressClass or a neov1alpha1.IngressClass.
type ingressClass struct {
	Name       string
	Controller string
	IsDefault  bool
}

const annotationDefaultIngressClass = "ingressclass.kubernetes.io/is-default-class"

// Supported ingress controller types.
const (
	ControllerTypeNginxOfficial  = "nginx.org/ingress-controller"
	ControllerTypeNginxCommunity = "k8s.io/ingress-nginx"
	ControllerTypeTraefik        = "traefik.io/ingress-controller"
)

// Watcher watches for IngressClass resources, maintaining a local cache of these resources,
// updated as they are created, modified or deleted.
// It watches for netv1.IngressClass, netv1beta1.IngressClass and neov1alpha1.IngressClass.
type Watcher struct {
	mu             sync.RWMutex
	ingressClasses map[ktypes.UID]ingressClass
}

// NewWatcher creates a new Watcher to track IngressClass resources.
func NewWatcher() *Watcher {
	return &Watcher{
		ingressClasses: make(map[ktypes.UID]ingressClass),
	}
}

// OnAdd implements Kubernetes cache.ResourceEventHandler so it can be used as an informer event handler.
func (w *Watcher) OnAdd(obj interface{}) {
	w.upsert(obj)
}

// OnUpdate implements Kubernetes cache.ResourceEventHandler so it can be used as an informer event handler.
func (w *Watcher) OnUpdate(_, newObj interface{}) {
	w.upsert(newObj)
}

// OnDelete implements Kubernetes cache.ResourceEventHandler so it can be used as an informer event handler.
func (w *Watcher) OnDelete(obj interface{}) {
	w.mu.Lock()
	defer w.mu.Unlock()

	switch v := obj.(type) {
	case *netv1.IngressClass:
		delete(w.ingressClasses, v.ObjectMeta.UID)
	case *netv1beta1.IngressClass:
		delete(w.ingressClasses, v.ObjectMeta.UID)
	case *neov1alpha1.IngressClass:
		delete(w.ingressClasses, v.ObjectMeta.UID)
	default:
		log.Error().
			Str("component", "ingressClassWatcher").
			Str("type", fmt.Sprintf("%T", obj)).
			Msg("Received delete event of unknown type")
	}
}

func (w *Watcher) upsert(obj interface{}) {
	w.mu.Lock()
	defer w.mu.Unlock()

	switch v := obj.(type) {
	case *netv1.IngressClass:
		w.ingressClasses[v.ObjectMeta.UID] = ingressClass{
			Name:       v.ObjectMeta.Name,
			Controller: v.Spec.Controller,
			IsDefault:  v.ObjectMeta.Annotations[annotationDefaultIngressClass] == "true",
		}
	case *netv1beta1.IngressClass:
		w.ingressClasses[v.ObjectMeta.UID] = ingressClass{
			Name:       v.ObjectMeta.Name,
			Controller: v.Spec.Controller,
			IsDefault:  v.ObjectMeta.Annotations[annotationDefaultIngressClass] == "true",
		}
	case *neov1alpha1.IngressClass:
		w.ingressClasses[v.ObjectMeta.UID] = ingressClass{
			Name:       v.ObjectMeta.Name,
			Controller: v.Spec.Controller,
			IsDefault:  v.ObjectMeta.Annotations[annotationDefaultIngressClass] == "true",
		}
	default:
		log.Error().
			Str("component", "ingressClassWatcher").
			Str("type", fmt.Sprintf("%T", obj)).
			Msg("Received upsert event of unknown type")
	}
}

// GetController returns the controller of the IngressClass matching the given name. If no IngressClass
// is found, an empty string is returned.
func (w *Watcher) GetController(name string) string {
	w.mu.RLock()
	defer w.mu.RUnlock()

	for _, class := range w.ingressClasses {
		if class.Name == name {
			return class.Controller
		}
	}

	return ""
}

// GetDefaultController returns the controller of the IngressClass that is noted as default.
// If no IngressClass is noted as default, an empty string is returned.
// If multiple IngressClasses are marked as default, an error is returned instead.
func (w *Watcher) GetDefaultController() (string, error) {
	w.mu.RLock()
	defer w.mu.RUnlock()

	var ctrlr string
	for _, ic := range w.ingressClasses {
		if ic.IsDefault {
			if ctrlr == "" {
				ctrlr = ic.Controller
				continue
			}
			return "", errors.New("multiple default ingress classes found")
		}
	}

	return ctrlr, nil
}
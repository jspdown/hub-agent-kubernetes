package state

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	traefikv1alpha1 "github.com/traefik/neo-agent/pkg/crd/api/traefik/v1alpha1"
	neokubemock "github.com/traefik/neo-agent/pkg/crd/generated/client/neo/clientset/versioned/fake"
	traefikkubemock "github.com/traefik/neo-agent/pkg/crd/generated/client/traefik/clientset/versioned/fake"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubemock "k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/kubernetes/scheme"
)

// Mandatory to be able to parse traefik.containo.us/v1alpha1 resources.
func init() {
	err := traefikv1alpha1.AddToScheme(scheme.Scheme)
	if err != nil {
		panic(err)
	}
}

func TestFetcher_GetIngressRoutes(t *testing.T) {
	tests := []struct {
		desc    string
		want    map[string]*IngressRoute
		fixture string
	}{
		{
			desc:    "One service",
			fixture: "ingress-route-one-service.yml",
			want: map[string]*IngressRoute{
				"name@ns.ingressroute.traefik.containo.us": {
					ResourceMeta: ResourceMeta{
						Kind:      ResourceKindIngressRoute,
						Group:     traefikv1alpha1.GroupName,
						Name:      "name",
						Namespace: "ns",
					},
					IngressMeta: IngressMeta{
						ClusterID:  "cluster-id",
						Controller: IngressControllerTypeTraefik,
					},
					TLS: &IngressRouteTLS{
						Domains: []traefikv1alpha1.Domain{
							{
								Main: "foo.com",
								SANs: []string{"bar.foo.com"},
							},
						},
						SecretName: "secret",
					},
					Routes: []Route{
						{
							Match: "Host(`foo.com`)",
							Services: []RouteService{
								{
									Name:       "service",
									Namespace:  "ns",
									PortNumber: 80,
								},
							},
						},
					},
					Services: []string{"service@ns"},
				},
			},
		},
		{
			desc:    "One service with an internal Traefik service",
			fixture: "ingress-route-one-internal-traefik-service.yml",
			want: map[string]*IngressRoute{
				"name@ns.ingressroute.traefik.containo.us": {
					ResourceMeta: ResourceMeta{
						Kind:      ResourceKindIngressRoute,
						Group:     traefikv1alpha1.GroupName,
						Name:      "name",
						Namespace: "ns",
					},
					IngressMeta: IngressMeta{
						ClusterID:  "cluster-id",
						Controller: IngressControllerTypeTraefik,
					},
					Routes: []Route{
						{
							Match: "Host(`api.localhost`)",
						},
					},
				},
			},
		},
		{
			desc:    "One Weighted Traefik service",
			fixture: "ingress-route-one-weighted-traefik-service.yml",
			want: map[string]*IngressRoute{
				"name@ns.ingressroute.traefik.containo.us": {
					ResourceMeta: ResourceMeta{
						Kind:      ResourceKindIngressRoute,
						Group:     traefikv1alpha1.GroupName,
						Name:      "name",
						Namespace: "ns",
					},
					IngressMeta: IngressMeta{
						ClusterID:  "cluster-id",
						Controller: IngressControllerTypeTraefik,
					},
					TLS: &IngressRouteTLS{
						Domains: []traefikv1alpha1.Domain{
							{
								Main: "foo.com",
								SANs: []string{"bar.foo.com"},
							},
						},
						SecretName: "secret",
					},
					Routes: []Route{
						{
							Match: "Host(`foo.com`)",
							Services: []RouteService{
								{
									Name:       "service1",
									Namespace:  "ns",
									PortNumber: 80,
								},
								{
									Name:       "service2",
									Namespace:  "ns",
									PortNumber: 80,
								},
							},
						},
					},
					Services: []string{"service1@ns", "service2@ns"},
				},
			},
		},
		{
			desc:    "One Mirroring Traefik service",
			fixture: "ingress-route-one-mirroring-traefik-service.yml",
			want: map[string]*IngressRoute{
				"name@ns.ingressroute.traefik.containo.us": {
					ResourceMeta: ResourceMeta{
						Kind:      ResourceKindIngressRoute,
						Group:     traefikv1alpha1.GroupName,
						Name:      "name",
						Namespace: "ns",
					},
					IngressMeta: IngressMeta{
						ClusterID:  "cluster-id",
						Controller: IngressControllerTypeTraefik,
					},
					TLS: &IngressRouteTLS{
						Domains: []traefikv1alpha1.Domain{
							{
								Main: "foo.com",
								SANs: []string{"bar.foo.com"},
							},
						},
						SecretName: "secret",
					},
					Routes: []Route{
						{
							Match: "Host(`foo.com`)",
							Services: []RouteService{
								{
									Name:       "service1",
									Namespace:  "ns2",
									PortNumber: 80,
								},
							},
						},
					},
					Services: []string{"service1@ns2"},
				},
			},
		},
		{
			desc:    "Two Weighted Traefik service",
			fixture: "ingress-route-two-weighted-traefik-service.yml",
			want: map[string]*IngressRoute{
				"name@ns.ingressroute.traefik.containo.us": {
					ResourceMeta: ResourceMeta{
						Kind:      ResourceKindIngressRoute,
						Group:     traefikv1alpha1.GroupName,
						Name:      "name",
						Namespace: "ns",
					},
					IngressMeta: IngressMeta{
						ClusterID:  "cluster-id",
						Controller: IngressControllerTypeTraefik,
					},
					TLS: &IngressRouteTLS{
						Domains: []traefikv1alpha1.Domain{
							{
								Main: "foo.com",
								SANs: []string{"bar.foo.com"},
							},
						},
						SecretName: "secret",
					},
					Routes: []Route{
						{
							Match: "Host(`foo.com`)",
							Services: []RouteService{
								{
									Name:       "service1",
									Namespace:  "ns",
									PortNumber: 80,
								},
								{
									Name:       "service2",
									Namespace:  "ns2",
									PortNumber: 80,
								},
								{
									Name:       "service3",
									Namespace:  "ns2",
									PortNumber: 80,
								},
							},
						},
					},
					Services: []string{"service1@ns", "service2@ns2", "service3@ns2"},
				},
			},
		},
		{
			desc:    "Two Mirroring Traefik service",
			fixture: "ingress-route-two-mirroring-traefik-service.yml",
			want: map[string]*IngressRoute{
				"name@ns.ingressroute.traefik.containo.us": {
					ResourceMeta: ResourceMeta{
						Kind:      ResourceKindIngressRoute,
						Group:     traefikv1alpha1.GroupName,
						Name:      "name",
						Namespace: "ns",
					},
					IngressMeta: IngressMeta{
						ClusterID:  "cluster-id",
						Controller: IngressControllerTypeTraefik,
					},
					Routes: []Route{
						{
							Match: "Host(`foo.com`)",
							Services: []RouteService{
								{
									Name:       "service",
									Namespace:  "ns2",
									PortNumber: 80,
								},
							},
						},
					},
					Services: []string{"service@ns2"},
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			objects := loadK8sObjects(t, "fixtures/ingress-route/"+test.fixture)

			kubeClient := kubemock.NewSimpleClientset()
			// Faking having Traefik CRDs installed on cluster.
			kubeClient.Resources = append(kubeClient.Resources, &metav1.APIResourceList{
				GroupVersion: traefikv1alpha1.SchemeGroupVersion.String(),
				APIResources: []metav1.APIResource{
					{
						Kind: ResourceKindIngressRoute,
					},
					{
						Kind: ResourceKindTraefikService,
					},
				},
			})

			neoClient := neokubemock.NewSimpleClientset()
			traefikClient := traefikkubemock.NewSimpleClientset(objects...)

			f, err := watchAll(context.Background(), kubeClient, neoClient, traefikClient, "v1.20.1", "cluster-id")
			require.NoError(t, err)

			got, err := f.getIngressRoutes("cluster-id")
			require.NoError(t, err)

			assert.Equal(t, test.want, got)
		})
	}
}
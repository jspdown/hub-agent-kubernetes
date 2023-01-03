/*
Copyright (C) 2022-2023 Traefik Labs

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.
*/

package edgeingress

import (
	"fmt"
	"strings"
	"time"

	hubv1alpha1 "github.com/traefik/hub-agent-kubernetes/pkg/crd/api/hub/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Certificate represents the certificate returned by the platform.
type Certificate struct {
	Certificate []byte `json:"certificate"`
	PrivateKey  []byte `json:"privateKey"`
}

// EdgeIngress is an ingress exposed on the edge.
type EdgeIngress struct {
	WorkspaceID string `json:"workspaceId"`
	ClusterID   string `json:"clusterId"`
	Namespace   string `json:"namespace"`
	Name        string `json:"name"`

	Domain        string         `json:"domain"`
	CustomDomains []CustomDomain `json:"customDomains"`

	Version string  `json:"version"`
	Service Service `json:"service"`
	ACP     *ACP    `json:"acp,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CustomDomain holds domain information.
type CustomDomain struct {
	Name     string `json:"name"`
	Verified bool   `json:"verified"`
}

// Service is a service used by the edge ingress.
type Service struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

// ACP is an ACP used by the edge ingress.
type ACP struct {
	Name string `json:"name"`
}

// Resource builds the v1alpha1 EdgeIngress resource.
func (e *EdgeIngress) Resource() (*hubv1alpha1.EdgeIngress, error) {
	var customDomain []string
	for _, domain := range e.CustomDomains {
		customDomain = append(customDomain, domain.Name)
	}

	spec := hubv1alpha1.EdgeIngressSpec{
		Service: hubv1alpha1.EdgeIngressService{
			Name: e.Service.Name,
			Port: e.Service.Port,
		},
		CustomDomains: customDomain,
	}

	if e.ACP != nil {
		spec.ACP = &hubv1alpha1.EdgeIngressACP{
			Name: e.ACP.Name,
		}
	}

	specHash, err := spec.Hash()
	if err != nil {
		return nil, fmt.Errorf("compute spec hash: %w", err)
	}

	var urls []string
	var verifiedCustomDomains []string
	for _, customDomain := range e.CustomDomains {
		if !customDomain.Verified {
			continue
		}

		urls = append(urls, "https://"+customDomain.Name)
		verifiedCustomDomains = append(verifiedCustomDomains, customDomain.Name)
	}

	urls = append(urls, "https://"+e.Domain)

	return &hubv1alpha1.EdgeIngress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      e.Name,
			Namespace: e.Namespace,
		},
		Spec: spec,
		Status: hubv1alpha1.EdgeIngressStatus{
			Version:       e.Version,
			SyncedAt:      metav1.Now(),
			Domain:        e.Domain,
			CustomDomains: verifiedCustomDomains,
			URLs:          strings.Join(urls, ","),
			Connection:    hubv1alpha1.EdgeIngressConnectionDown,
			SpecHash:      specHash,
		},
	}, nil
}

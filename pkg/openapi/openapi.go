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

package openapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/hashicorp/go-version"
	corev1 "k8s.io/api/core/v1"
)

const (
	annotationOpenAPIPath = "hub.traefik.io/openapi-path"
	annotationOpenAPIPort = "hub.traefik.io/openapi-port"
)

// Location describes the location of an OpenAPI specification.
type Location struct {
	Path string `json:"path"`
	Port int    `json:"port"`
}

// GetLocationFromService retrieve the location of an OpenAPI specification on the given service.
func GetLocationFromService(service *corev1.Service) (*Location, error) {
	aosPath, ok := service.Annotations[annotationOpenAPIPath]
	if !ok {
		return nil, nil
	}

	var portStr string
	portStr, ok = service.Annotations[annotationOpenAPIPort]
	if !ok {
		return nil, nil
	}

	aosPort, err := strconv.ParseInt(portStr, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("%q must be a valid port", annotationOpenAPIPort)
	}

	var portFound bool
	for _, servicePort := range service.Spec.Ports {
		if int64(servicePort.Port) == aosPort {
			portFound = true
			break
		}
	}
	if !portFound {
		return nil, fmt.Errorf("%q contains a port which is not defined on the service", annotationOpenAPIPort)
	}

	return &Location{
		Path: aosPath,
		Port: int(aosPort),
	}, nil
}

// Loader loads OpenAPI Specifications.
type Loader struct{}

// LoadFromURI loads the OpenAPI Specification located at the given URL.
func (l *Loader) LoadFromURI(uri *url.URL) (*Spec, error) {
	// Create a new loader each time. Indeed, the openapi3 package caches the loaded specification documents. A change
	// on the document without a change on the uri would not be detected. If caching is needed it must be built on top
	// of this loader. Also, the loader wouldn't be safe for concurrent use otherwise.
	loader := openapi3.NewLoader()

	spec, err := loader.LoadFromURI(uri)
	if err != nil {
		return nil, err
	}

	// Spec needs to be wrapped, additional validations are required.
	return &Spec{
		spec: spec,
	}, nil
}

// Spec is an OpenAPI Specification.
type Spec struct {
	spec *openapi3.T
}

// UnmarshalJSON unmarshals the given bytes into itself.
func (s *Spec) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &s.spec)
}

// Validate validates the specification.
func (s *Spec) Validate(ctx context.Context) error {
	v, err := version.NewVersion(s.spec.OpenAPI)
	if err != nil {
		return fmt.Errorf("invalid version: %w", err)
	}

	major := v.Segments()[0]
	if major != 3 {
		return fmt.Errorf("unsupported version %q", v.String())
	}

	return s.spec.Validate(ctx)
}

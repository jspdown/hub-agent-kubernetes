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

package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubemock "k8s.io/client-go/kubernetes/fake"
)

func TestReadKey(t *testing.T) {
	cliCtx := &cli.Context{Context: context.Background()}
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "hub-secret",
			Namespace: "default",
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{
			"key": []byte("my-key"),
		},
	}

	clientSetHub := kubemock.NewSimpleClientset(secret)

	key, err := readKey(cliCtx, clientSetHub)
	require.NoError(t, err)
	require.Equal(t, "5e78863ed1ffb9fc66b1d61634b126bf", key)
}

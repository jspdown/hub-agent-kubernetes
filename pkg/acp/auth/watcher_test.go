package auth

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	neov1alpha1 "github.com/traefik/neo-agent/pkg/crd/api/neo/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func createPolicy(uid, name, ns string) *neov1alpha1.AccessControlPolicy {
	return &neov1alpha1.AccessControlPolicy{
		ObjectMeta: metav1.ObjectMeta{UID: types.UID(uid), Name: name, Namespace: ns},
		Spec: neov1alpha1.AccessControlPolicySpec{
			JWT: &neov1alpha1.AccessControlPolicyJWT{
				SigningSecret: "secret",
			},
		},
	}
}

func TestWatcher_OnAdd(t *testing.T) {
	switcher := NewHandlerSwitcher()
	watcher := NewWatcher(switcher)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	t.Cleanup(cancel)

	go watcher.Run(ctx)

	watcher.OnAdd(createPolicy("1", "my-policy-1", "test"))
	watcher.OnAdd(createPolicy("2", "my-policy-2", "test"))
	watcher.OnAdd(createPolicy("3", "my-policy-3", "foo"))

	time.Sleep(10 * time.Millisecond)

	testCases := []struct {
		desc     string
		path     string
		expected int
	}{
		{
			desc:     "my-policy-1",
			path:     "/my-policy-1@test",
			expected: http.StatusUnauthorized,
		},
		{
			desc:     "my-policy-2",
			path:     "/my-policy-2@test",
			expected: http.StatusUnauthorized,
		},
		{
			desc:     "my-policy-3",
			path:     "/my-policy-3@foo",
			expected: http.StatusUnauthorized,
		},
		{
			desc:     "unknown resource",
			path:     "/my-policy@test",
			expected: http.StatusNotFound,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "http://localhost"+test.path, nil)

			switcher.ServeHTTP(rw, req)

			assert.Equal(t, test.expected, rw.Code)
		})
	}
}

func TestWatcher_OnUpdate(t *testing.T) {
	switcher := NewHandlerSwitcher()
	watcher := NewWatcher(switcher)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	t.Cleanup(cancel)

	go watcher.Run(ctx)

	watcher.OnAdd(createPolicy("1", "my-policy-1", "test"))
	watcher.OnAdd(createPolicy("2", "my-policy-2", "test"))
	watcher.OnAdd(createPolicy("3", "my-policy-3", "foo"))

	watcher.OnUpdate(nil, createPolicy("1", "my-policy-1", "test"))
	watcher.OnUpdate(nil, createPolicy("2", "my-policy-2", "test"))
	watcher.OnUpdate(nil, createPolicy("3", "my-policy-3", "foo"))

	time.Sleep(10 * time.Millisecond)

	testCases := []struct {
		desc     string
		path     string
		expected int
	}{
		{
			desc:     "my-policy-1",
			path:     "/my-policy-1@test",
			expected: http.StatusUnauthorized,
		},
		{
			desc:     "my-policy-2",
			path:     "/my-policy-2@test",
			expected: http.StatusUnauthorized,
		},
		{
			desc:     "my-policy-3",
			path:     "/my-policy-3@foo",
			expected: http.StatusUnauthorized,
		},
		{
			desc:     "unknown resource",
			path:     "/my-policy@test",
			expected: http.StatusNotFound,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "http://localhost"+test.path, nil)

			switcher.ServeHTTP(rw, req)

			assert.Equal(t, test.expected, rw.Code)
		})
	}
}

func TestWatcher_OnDelete(t *testing.T) {
	switcher := NewHandlerSwitcher()
	watcher := NewWatcher(switcher)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	t.Cleanup(cancel)

	go watcher.Run(ctx)

	watcher.OnAdd(createPolicy("1", "my-policy-1", "test"))
	watcher.OnAdd(createPolicy("2", "my-policy-2", "test"))
	watcher.OnAdd(createPolicy("3", "my-policy-3", "foo"))

	watcher.OnDelete(createPolicy("1", "my-policy-1", "test"))
	watcher.OnDelete(createPolicy("2", "my-policy-2", "test"))
	watcher.OnDelete(createPolicy("3", "my-policy-3", "foo"))

	time.Sleep(10 * time.Millisecond)

	testCases := []struct {
		desc     string
		path     string
		expected int
	}{
		{
			desc:     "my-policy-1",
			path:     "/my-policy-1@test",
			expected: http.StatusNotFound,
		},
		{
			desc:     "my-policy-2",
			path:     "/my-policy-2@test",
			expected: http.StatusNotFound,
		},
		{
			desc:     "my-policy-3",
			path:     "/my-policy-3@foo",
			expected: http.StatusNotFound,
		},
		{
			desc:     "unknown resource",
			path:     "/my-policy@test",
			expected: http.StatusNotFound,
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			rw := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "http://localhost"+test.path, nil)

			switcher.ServeHTTP(rw, req)

			assert.Equal(t, test.expected, rw.Code)
		})
	}
}
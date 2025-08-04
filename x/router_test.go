// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package x

import (
	"net/http"
	"testing"

	"github.com/gobuffalo/httptest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRouterAdmin(t *testing.T) {
	require.NotEmpty(t, NewRouterAdmin())
	require.NotEmpty(t, NewRouterPublic())
}

func TestCacheHandling(t *testing.T) {
	router := NewRouterPublic()
	ts := httptest.NewServer(router)
	t.Cleanup(ts.Close)

	router.HandleFunc("GET /foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	router.HandleFunc("DELETE /foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	router.HandleFunc("POST /foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	router.HandleFunc("PUT /foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	router.HandleFunc("PATCH /foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	for _, method := range []string{} {
		req, _ := http.NewRequest(method, ts.URL+"/foo", nil)
		res, err := ts.Client().Do(req)
		require.NoError(t, err)
		assert.EqualValues(t, "0", res.Header.Get("Cache-Control"))
	}
}

func TestAdminPrefix(t *testing.T) {
	router := NewRouterAdmin()
	ts := httptest.NewServer(router)
	t.Cleanup(ts.Close)

	router.HandleFunc("GET /foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	router.HandleFunc("DELETE /foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	router.HandleFunc("POST /foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	router.HandleFunc("PUT /foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	router.HandleFunc("PATCH /foo", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	for _, method := range []string{} {
		req, _ := http.NewRequest(method, ts.URL+"/admin/foo", nil)
		res, err := ts.Client().Do(req)
		require.NoError(t, err)
		assert.EqualValues(t, http.StatusNoContent, res.StatusCode)
	}
}

package router

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRouter(t *testing.T) {
	v1Called := false
	showUserCalled := false

	router := &Router{
		routes: routeMethodMap{
			http.MethodGet: routeList{
				route{pathSegments: []string{"api", "v1"}, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					v1Called = true
					assert.Equal(t, "/api/v1", r.URL.Path)
					w.WriteHeader(http.StatusTeapot)
				})},
				route{pathSegments: []string{"api", "v1", "users", ":user_id"}, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					showUserCalled = true
					assert.Equal(t, "/api/v1/users/10", r.URL.Path)
					assert.Equal(t, "10", PathVar(r, "user_id"))
					w.WriteHeader(http.StatusOK)
				})},
			},
		},
	}

	server := httptest.NewServer(router)
	defer server.Close()

	res, err := http.Get(fmt.Sprintf("%s/api/v1", server.URL))
	require.NoError(t, err)
	assert.Equal(t, http.StatusTeapot, res.StatusCode)

	res, err = http.Get(fmt.Sprintf("%s/api/v1/users/10", server.URL))
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	res, err = http.Get(fmt.Sprintf("%s/api/v1/users", server.URL))
	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)

	assert.True(t, v1Called)
	assert.True(t, showUserCalled)
}

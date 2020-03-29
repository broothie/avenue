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
	eventQueryCalled := false

	router := &Router{
		routes: routeMap{
			http.MethodGet: routes{
				route{segments: []string{"api", "v1"}, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					v1Called = true
					assert.Equal(t, "/api/v1", r.URL.Path)
					w.WriteHeader(http.StatusTeapot)
				})},
				route{segments: []string{"api", "v1", "users", ":user_id"}, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					showUserCalled = true
					assert.Equal(t, "/api/v1/users/10", r.URL.Path)
					assert.Equal(t, "10", PathVar(r, "user_id"))
					w.WriteHeader(http.StatusOK)
				})},
			},
			http.MethodPost: routes{
				route{segments: []string{"api", "v1", "events"}, queries: map[string]string{"page": ""}, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					eventQueryCalled = true
					assert.Equal(t, "/api/v1/events", r.URL.Path)
					assert.Equal(t, "5", r.URL.Query().Get("page"))
					w.WriteHeader(http.StatusAccepted)
				})},
				route{segments: []string{"api", "v1", "events"}, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, "/api/v1/events", r.URL.Path)
					w.WriteHeader(http.StatusOK)
					assert.Fail(t, "events w/o query should not be called")
				})},
				route{segments: []string{"api", "v1", "posts"}, queries: map[string]string{"name": "andrew"}, handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					assert.Equal(t, "/api/v1/posts", r.URL.Path)
					w.WriteHeader(http.StatusOK)
					assert.Fail(t, "posts w/o query value should not be called")
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

	res, err = http.Post(fmt.Sprintf("%s/api/v1/events?page=5", server.URL), "", nil)
	require.NoError(t, err)
	assert.Equal(t, http.StatusAccepted, res.StatusCode)

	res, err = http.Post(fmt.Sprintf("%s/api/v1/posts?name=todd", server.URL), "", nil)
	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)

	res, err = http.Get(fmt.Sprintf("%s/api/v1/users", server.URL))
	require.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, res.StatusCode)

	assert.True(t, v1Called)
	assert.True(t, showUserCalled)
	assert.True(t, eventQueryCalled)
}

func TestRouter_AddRoute(t *testing.T) {
	router := New()
	router.AddRoute(http.MethodPost, "api/:version", nil, nil, nil)
	router.AddRoute(http.MethodPost, "api", nil, nil, nil)
	router.AddRoute(http.MethodPost, "api/v1", nil, nil, nil)

	assert.Equal(t, router.routes[http.MethodPost][0].segments, []string{"api"})
	assert.Equal(t, router.routes[http.MethodPost][1].segments, []string{"api", "v1"})
	assert.Equal(t, router.routes[http.MethodPost][2].segments, []string{"api", ":version"})
}

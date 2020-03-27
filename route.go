package drr

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	method  string
	path    string
	queries []Pair
	headers []Pair
	handler http.Handler

	summary string

	parent *Route
	routes []*Route
	router *mux.Router
}

type Pair struct {
	Key      string
	Value    string
	Required bool
}

func New(path string) *Route {
	return &Route{
		path:    path,
		queries: make([]Pair, 0),
		headers: make([]Pair, 0),
		routes:  make([]*Route, 0),
		router:  mux.NewRouter(),
	}
}

func (r *Route) newChild() *Route {
	child := r.copy()
	child.parent = r
	return child
}

func (r *Route) copy() *Route {
	queries := make([]Pair, len(r.queries))
	copy(queries, r.queries)

	headers := make([]Pair, len(r.headers))
	copy(headers, r.headers)

	routes := make([]*Route, len(r.routes))
	copy(routes, r.routes)

	newRoute := new(Route)
	*newRoute = *r
	newRoute.queries = queries
	newRoute.headers = headers
	newRoute.routes = routes
	return newRoute
}

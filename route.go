package ave

import (
	"net/http"

	"github.com/broothie/avenue/internal/router"
)

type (
	Route struct {
		method      string
		path        string
		queries     []Query
		headers     []Header
		handler     http.Handler
		middlewares []func(http.Handler) http.Handler

		documentation Documentation

		parent    *Route
		endpoints []*Route
		router    router.Router
		//router    *mux.Router
	}

	Pair struct {
		Name     string
		Value    string
		Type     string
		Required bool
	}

	Query  Pair
	Header Pair

	Queries []Query
	Headers []Header
)

func New(path string) *Route {
	return &Route{
		path:   path,
		router: router.New(),
	}
}

func Root() *Route {
	return New("/")
}

func (r *Route) Root() *Route {
	root := r
	for root.parent != nil {
		root = root.parent
	}

	return root
}

func (r *Route) newChild() *Route {
	child := r.copy()
	child.parent = r
	return child
}

func (r *Route) copy() *Route {
	newRoute := new(Route)
	*newRoute = *r
	newRoute.documentation.Summary = ""
	newRoute.documentation.Description = ""

	queries := make([]Query, len(r.queries))
	copy(queries, r.queries)
	newRoute.queries = queries

	headers := make([]Header, len(r.headers))
	copy(headers, r.headers)
	newRoute.headers = headers

	middlewares := make([]func(http.Handler) http.Handler, len(r.middlewares))
	copy(middlewares, r.middlewares)
	newRoute.middlewares = middlewares

	body := make([]Key, len(r.documentation.Body))
	copy(body, r.documentation.Body)
	newRoute.documentation.Body = body

	responses := make([]Response, len(r.documentation.Responses))
	copy(responses, r.documentation.Responses)
	newRoute.documentation.Responses = responses

	endpoints := make([]*Route, len(r.endpoints))
	copy(endpoints, r.endpoints)
	newRoute.endpoints = endpoints

	return newRoute
}

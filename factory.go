package drr

import (
	"net/http"
	gopath "path"
)

func (r *Route) Method(method string) *Route {
	route := r.newChild()
	route.method = method
	return route
}

func (r *Route) Path(path string) *Route {
	route := r.newChild()
	route.path = path
	return route
}

func (r *Route) Nest(path string, f func(route *Route)) {
	route := r.newChild()
	route.path = gopath.Join(r.path, path)
	f(route)
}

func (r *Route) Queries(queries ...Pair) *Route {
	route := r.newChild()
	route.queries = append(route.queries, queries...)
	return route
}

func (r *Route) Headers(headers ...Pair) *Route {
	route := r.newChild()
	route.headers = append(route.headers, headers...)
	return route
}

func (r *Route) Middleware(middleware ...func(http.Handler) http.Handler) *Route {
	route := r.newChild()
	route.middlewares = append(route.middlewares, middleware...)
	return route
}

func (r *Route) Summary(summary string) *Route {
	route := r.newChild()
	route.summary = summary
	return route
}

func (r *Route) Description(description string) *Route {
	route := r.newChild()
	route.description = description
	return route
}

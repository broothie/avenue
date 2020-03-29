package drr

import (
	"net/http"
	gopath "path"
)

func (r *Route) Nest(f func(route *Route)) {
	f(r)
}

func (r *Route) Method(method string) *Route {
	route := r.newChild()
	route.method = method
	return route
}

func (r *Route) Path(path string) *Route {
	route := r.newChild()
	route.path = gopath.Join(r.path, path)
	return route
}

func (r *Route) Queries(queries ...Query) *Route {
	route := r.newChild()
	route.queries = append(route.queries, queries...)
	return route
}

func (r *Route) Headers(headers ...Header) *Route {
	route := r.newChild()
	route.headers = append(route.headers, headers...)
	return route
}

func (r *Route) Middleware(middleware ...func(http.Handler) http.Handler) *Route {
	route := r.newChild()
	route.middlewares = append(route.middlewares, middleware...)
	return route
}

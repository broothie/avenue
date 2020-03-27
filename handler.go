package drr

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/mux"
)

func (r *Route) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.router.ServeHTTP(writer, request)
}

//go:generate go run generate_method_funcs.go -tags=generate
func (r *Route) Handler(handler http.Handler) {
	endpoint := r.newChild()
	if r.method == "" {
		r.method = http.MethodGet
	}

	endpoint.handler = handler
	r.addRoute(endpoint)
}

func (r *Route) HandlerFunc(handlerFunc http.HandlerFunc) {
	r.Handler(handlerFunc)
}

func (r *Route) addRoute(route *Route) {
	r.routes = append(r.routes, route)

	sort.SliceStable(r.routes, func(i, j int) bool {
		leftRoute, rightRoute := r.routes[i], r.routes[j]
		leftSegments, rightSegments := strings.Split(leftRoute.path, "/"), strings.Split(rightRoute.path, "/")
		if len(leftSegments) != len(rightSegments) {
			return len(leftSegments) < len(rightSegments)
		}

		leftLast, rightLast := leftSegments[len(leftSegments)-1], rightSegments[len(rightSegments)-1]
		if leftLast[0] == '{' && rightLast[0] != '{' {
			return false
		}

		return len(leftLast) < len(rightLast)
	})

	router := mux.NewRouter()
	for _, route := range r.routes {
		var queries []string
		for _, pair := range route.queries {
			if pair.Required {
				value := pair.Value
				if value == "" {
					value = fmt.Sprintf("{%s}", pair.Key)
				}

				queries = append(queries, pair.Key, value)
			}
		}

		var headers []string
		for _, pair := range route.headers {
			if pair.Required {
				value := pair.Value
				if pair.Value == "" {
					value = ".*"
				}

				headers = append(headers, pair.Key, value)
			}
		}

		router.
			Methods(route.method).
			Path(route.path).
			Queries(queries...).
			HeadersRegexp(headers...).
			Handler(applyMiddlewares(route.handler, route.middlewares...))
	}

	r.router = router
	if r.parent != nil {
		r.parent.addRoute(route)
	}
}

func applyMiddlewares(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}

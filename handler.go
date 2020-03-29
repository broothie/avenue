package ave

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"

	"github.com/gorilla/mux"
)

func (r *Route) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	r.router.ServeHTTP(writer, request)
}

func (r *Route) Handler(handler http.Handler) {
	endpoint := r.newChild()
	if r.method == "" {
		r.method = http.MethodGet
	}

	endpoint.handler = handler

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go r.addEndpoint(endpoint, wg)
	wg.Wait()
}

func (r *Route) HandlerFunc(handlerFunc http.HandlerFunc) {
	r.Handler(handlerFunc)
}

func (r *Route) addEndpoint(endpoint *Route, wg *sync.WaitGroup) {
	defer wg.Done()
	if r.parent != nil {
		wg.Add(1)
		go r.parent.addEndpoint(endpoint, wg)
	}

	r.endpoints = append(r.endpoints, endpoint)
	r.router = buildRouter(r.endpoints)
}

func buildRouter(endpoints []*Route) *mux.Router {
	sort.SliceStable(sortEndpoints(endpoints))

	router := mux.NewRouter()
	for _, endpoint := range endpoints {
		var queries []string
		for _, pair := range endpoint.queries {
			if pair.Required {
				value := pair.Value
				if value == "" {
					value = fmt.Sprintf("{%s}", pair.Name)
				}

				queries = append(queries, pair.Name, value)
			}
		}

		var headers []string
		for _, pair := range endpoint.headers {
			if pair.Required {
				value := pair.Value
				if pair.Value == "" {
					value = ".*"
				}

				headers = append(headers, pair.Name, value)
			}
		}

		router.
			Methods(endpoint.method).
			Path(endpoint.path).
			Queries(queries...).
			HeadersRegexp(headers...).
			Handler(applyMiddlewares(endpoint.handler, endpoint.middlewares...))
	}

	return router
}

func sortEndpoints(endpoints []*Route) ([]*Route, func(i, j int) bool) {
	return endpoints, func(i, j int) bool {
		left, right := endpoints[i], endpoints[j]
		lSegs, rSegs := strings.Split(left.path, "/"), strings.Split(right.path, "/")
		if len(lSegs) != len(rSegs) {
			return len(lSegs) < len(rSegs)
		}

		lLast, rLast := lSegs[len(lSegs)-1], rSegs[len(rSegs)-1]
		if lLast[0] == '{' && rLast[0] != '{' {
			return false
		}

		return len(lLast) < len(rLast)
	}
}

func applyMiddlewares(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}

package router

import (
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	routes routeMethodMap
}

type routeMethodMap map[string]routeList

type routeList []route

type route struct {
	pathSegments []string
	handler      http.Handler
}

type query struct {
	key   string
	value string
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	routes, ok := r.routes[request.Method]
	if !ok {
		r.notFound(writer)
		return
	}

	for _, route := range routes {
		if match, request := route.match(request); match {
			route.handler.ServeHTTP(writer, request)
			return
		}
	}

	r.notFound(writer)
}

func (r *Router) notFound(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusNotFound)
	fmt.Fprint(writer, "route not found")
}

func (r route) match(request *http.Request) (bool, *http.Request) {
	return r.matchPath(request)
}

func (r route) matchPath(request *http.Request) (bool, *http.Request) {
	reqSegments := strings.Split(request.URL.Path, "/")[1:] // [1:] to skip the first `""` from the leading `/`
	rteSegments := r.pathSegments
	if len(reqSegments) != len(rteSegments) {
		return false, nil
	}

	for i, reqSegment := range reqSegments {
		rteSegment := r.pathSegments[i]

		// If not a named segment, simply check for match
		if !strings.HasPrefix(rteSegment, ":") {
			if reqSegment != rteSegment {
				return false, nil
			} else {
				continue
			}
		}

		// Named segment can't have empty value
		if reqSegment == "" {
			return false, nil
		}

		request = setPathVarOnRequest(request, strings.TrimPrefix(rteSegment, ":"), reqSegment)
	}

	return true, request
}

func (r route) matchQuery(request *http.Request) bool {

}

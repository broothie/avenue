package router

import (
	"net/http"
	gopath "path"
	"sort"
	"strings"
)

type Router struct {
	routes routeMap
}

func New() Router {
	return Router{
		routes: make(routeMap),
	}
}

type routeMap map[string]routes

type routes []route

type route struct {
	segments []string
	queries  map[string]string
	headers  map[string]string
	handler  http.Handler
}

func (r Router) AddRoute(method, path string, queries, headers map[string]string, handler http.Handler) {
	rte := route{
		segments: strings.Split(gopath.Join("/", path), "/")[1:],
		queries:  queries,
		headers:  headers,
		handler:  handler,
	}

	method = strings.ToUpper(method)
	if _, ok := r.routes[method]; !ok {
		r.routes[method] = routes{rte}
		return
	}

	r.routes[method] = append(r.routes[method], rte)
	sort.SliceStable(routeSorter(r.routes[method]))
}

func routeSorter(routes routes) (interface{}, func(int, int) bool) {
	return routes, func(i, j int) bool {
		lSegs, rSegs := routes[i].segments, routes[j].segments
		for i := 0; i < max(len(lSegs), len(rSegs)); i++ {
			lSeg, rSeg := safeGet(lSegs, i), safeGet(rSegs, i)
			if lSeg == nil {
				return true
			}

			if rSeg == nil {
				return false
			}

			if strings.HasPrefix(*lSeg, ":") {
				return false
			}

			if strings.HasPrefix(*rSeg, ":") {
				return true
			}
		}

		return true
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func safeGet(strings []string, i int) *string {
	if i >= len(strings) {
		return nil
	}

	return &strings[i]
}

package router

import (
	"fmt"
	"net/http"
)

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

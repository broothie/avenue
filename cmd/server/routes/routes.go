package routes

import (
	"fmt"
	"net/http"
	"time"

	drr "github.com/broothie/drroute"
)

func Routes() *drr.Route {
	route := drr.Root()

	route.
		DocOmit().
		DocSummary("file server").
		Method(http.MethodGet).
		Handler(http.FileServer(http.Dir(".")))

	route.Path("api").Nest(func(route *drr.Route) {
		route.Path("v1").Nest(func(route *drr.Route) {
			route.Path("users").Nest(func(route *drr.Route) {
				route = route.Middleware(timeMiddleware)

				route.
					DocSummary("index users").
					DocDescription("returns a list of users").
					Method(http.MethodGet).
					Queries(drr.Query{Name: "page", Type: "integer", Required: true}).
					HandlerFunc(indexUsers)

				route.
					DocSummary("create user").
					DocDescription("creates a user").
					Method(http.MethodPost).
					HandlerFunc(indexUsers)

				route.Path("{user_id}").Nest(func(route *drr.Route) {
					route.
						DocSummary("show user").
						Method(http.MethodGet).
						DocResponses(drr.Response{Status: http.StatusOK}).
						HandlerFunc(indexUsers)
				})
			})

			route.
				DocSummary("create event").
				DocDescription("creates an event and returns its details").
				Method(http.MethodPost).
				Path("events").
				DocBody(drr.Key{Name: "title", Type: "string", Required: true}).
				HandlerFunc(createEvent)
		})
	})

	return route
}

func indexUsers(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `[]`)
}

func createEvent(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{}`)
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("time", time.Now().String())
		next.ServeHTTP(w, r)
	})
}

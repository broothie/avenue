package routes

import (
	"fmt"
	"net/http"
	"time"

	ave "github.com/broothie/avenue"
	"github.com/broothie/avenue/openapi"
	"github.com/go-openapi/runtime/middleware"
)

var redocOpts = middleware.RedocOpts{
	BasePath: "/api/v1",
	SpecURL:  "/api/v1/docs/specs/openapi.yml",
}

func Routes() *ave.Route {
	route := ave.Root()

	route.
		DocSummary("file server").
		Method(http.MethodGet).
		Path("/static/{_:.*}").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("."))))

	route.Path("api").Nest(func(route *ave.Route) {
		route.Path("v1").Nest(func(route *ave.Route) {
			route.Path("docs").Nest(func(route *ave.Route) {
				route.
					DocSummary("api docs").
					Method(http.MethodGet).
					Handler(middleware.Redoc(redocOpts, http.NotFoundHandler()))

				route.
					DocSummary("OpenAPI spec").
					Method(http.MethodGet).
					Path("/specs/{_:.*}").
					Handler(openapi.SpecHandler(route.Root()))
			})

			route.Path("users").Nest(func(route *ave.Route) {
				route = route.Middleware(timeMiddleware)

				route.
					DocSummary("index users").
					DocDescription("returns a list of users").
					Method(http.MethodGet).
					Queries(ave.Query{Name: "page", Type: "integer", Required: true}).
					HandlerFunc(indexUsers)

				route.
					DocSummary("create user").
					DocDescription("creates a user").
					Method(http.MethodPost).
					HandlerFunc(indexUsers)

				route.Path("{user_id}").Nest(func(route *ave.Route) {
					route.
						DocSummary("show user").
						Method(http.MethodGet).
						DocResponses(ave.Response{Status: http.StatusOK}).
						HandlerFunc(indexUsers)
				})
			})

			route.
				DocSummary("create event").
				DocDescription("creates an event and returns its details").
				Method(http.MethodPost).
				Path("events").
				DocBody(ave.Key{Name: "title", Type: "string", Required: true}).
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

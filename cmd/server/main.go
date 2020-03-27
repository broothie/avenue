package main

import (
	"fmt"
	"log"
	"net/http"

	drr "github.com/broothie/drroute"
)

func main() {
	api := drr.New("/api/v1")

	api.Nest("/users", func(users *drr.Route) {
		users.
			Summary("user index").
			Method(http.MethodGet).
			Queries(drr.Pair{Key: "page", Required: true}).
			Headers(drr.Pair{Key: "accept"}).
			HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusTeapot)
				fmt.Fprint(w, r.URL.Query().Get("page"))
			})

		users.Nest("/{user_id}", func(user *drr.Route) {
			user.
				Summary("user show").
				Method(http.MethodGet).
				Queries(drr.Pair{Key: "admin"}).
				Handler(stringHandler(http.StatusOK, "user show"))
		})

		users.Nest("/events", func(events *drr.Route) {
			events.Get(stringHandler(http.StatusAccepted, "event show"))
			events.Post(stringHandler(http.StatusAlreadyReported, "event create"))
		})
	})

	if err := api.GenerateDoc(); err != nil {
		log.Panic(err)
	}

	log.Panic(http.ListenAndServe(":8080", api))
}

func stringHandler(code int, s string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(code)
		fmt.Fprint(w, s)
	}
}

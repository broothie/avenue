package main

import (
	"log"
	"net/http"

	"github.com/broothie/drroute/cmd/server/routes"
	"github.com/broothie/drroute/openapi"
	"github.com/broothie/drroute/printer"
)

func main() {
	api := routes.Routes()

	printer.Print(api)
	if err := openapi.Generate(api); err != nil {
		log.Panic(err)
	}

	log.Panic(http.ListenAndServe(":8080", api))
}

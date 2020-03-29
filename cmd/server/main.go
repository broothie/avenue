package main

import (
	"log"
	"net/http"

	"github.com/broothie/avenue/cmd/server/routes"
	"github.com/broothie/avenue/openapi"
	"github.com/broothie/avenue/printer"
)

func main() {
	api := routes.Routes()

	printer.Print(api)
	if err := openapi.GenerateFile(api); err != nil {
		log.Panic(err)
	}

	log.Panic(http.ListenAndServe(":8080", api))
}

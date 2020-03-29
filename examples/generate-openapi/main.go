package main

import (
	"fmt"
	"os"

	"github.com/broothie/drroute/examples"

	"github.com/broothie/drroute/openapi"
)

func main() {
	route := examples.RouteFunc()
	if err := openapi.Generate(route); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/broothie/avenue/examples"
	"github.com/broothie/avenue/openapi"
)

func main() {
	route := examples.RouteFunc()
	if err := openapi.GenerateFile(route); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

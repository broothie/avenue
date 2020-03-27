package main

import (
	"fmt"
	"os"

	drr "github.com/broothie/drroute"
)

func main() {
	route := RouteFunc()
	if err := route.GenerateDoc(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func RouteFunc() *drr.Route {
	return drr.New("/")
}

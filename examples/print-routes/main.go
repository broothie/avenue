package main

import (
	"github.com/broothie/drroute/examples"
	"github.com/broothie/drroute/printer"
)

func main() {
	route := examples.RouteFunc()
	printer.Print(route)
}

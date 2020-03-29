package main

import (
	"github.com/broothie/avenue/examples"
	"github.com/broothie/avenue/printer"
)

func main() {
	route := examples.RouteFunc()
	printer.Print(route)
}

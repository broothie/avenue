package main

import "flag"

const importFileFmt = `
package main

import (
	"%s"
)

func main() {
	route := %s()
	route.GenerateDoc()
}
`

func main() {
	funcName := flag.String("func", "", "function which returns a *drr.Route")
	flag.Parse()

}

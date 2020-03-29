package main

import (
	"fmt"
	"os"
)

const importFileFmt = `
package main

import (
	"os"
	"%s"
)

func main() {
	if err := openapi.Generate(%s()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
`

func main() {
	if err := os.MkdirAll(fmt.Sprintf("%s/main.go"), os.ModePerm); err != nil {

	}
}

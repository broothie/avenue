//+build generate

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var methods = []string{
	http.MethodGet,
	http.MethodHead,
	http.MethodPost,
	http.MethodPut,
	http.MethodPatch,
	http.MethodDelete,
	http.MethodConnect,
	http.MethodOptions,
	http.MethodTrace,
}

func main() {
	builder := new(strings.Builder)

	builder.WriteString("package drr\n")
	builder.WriteString("\n")
	builder.WriteString("import \"net/http\"\n")
	builder.WriteString("\n")

	for _, method := range methods {
		camelCase := strings.Title(strings.ToLower(method))
		builder.WriteString(fmt.Sprintf("func (r *Route) %s(handler http.Handler) {\n", camelCase))
		builder.WriteString(fmt.Sprintf("	r.method = http.Method%s\n", camelCase))
		builder.WriteString("	r.Handler(handler)\n")
		builder.WriteString("}\n")
		builder.WriteString("\n")
	}

	if err := ioutil.WriteFile("method_funcs.go.txt", []byte(builder.String()), os.ModePerm); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

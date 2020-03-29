package printer

import (
	"fmt"
	"strconv"
	"strings"

	drr "github.com/broothie/drroute"
)

func Print(route *drr.Route) {
	fmt.Print(String(route))
}

func String(route *drr.Route) string {
	builder := new(strings.Builder)

	endpoints := route.EndpointInfo()
	maxMethodLength := maxMethodLength(endpoints)
	for _, endpoint := range endpoints {
		builder.WriteString(fmt.Sprintf("%-"+strconv.Itoa(maxMethodLength)+"s %s", endpoint.Method, endpoint.Path))

		if len(endpoint.Queries) > 0 {
			var queries []string
			for _, query := range endpoint.Queries {
				queryString := query.Name
				if !query.Required {
					queryString += "?"
				}

				queryString += "="
				if query.Value != "" {
					queryString += query.Value
				}

				queries = append(queries, queryString)
				builder.WriteString(fmt.Sprintf("?%s", strings.Join(queries, "&")))
			}
		}

		builder.WriteString("\n")
	}

	return builder.String()
}

func maxMethodLength(endpoints []drr.RouteInfo) int {
	max := 0
	for _, endpoint := range endpoints {
		if len(endpoint.Method) > max {
			max = len(endpoint.Method)
		}
	}

	return max
}

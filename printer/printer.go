package printer

import (
	"fmt"
	"strconv"
	"strings"

	ave "github.com/broothie/avenue"
)

func Print(route *ave.Route) {
	fmt.Print(String(route))
}

func String(route *ave.Route) string {
	builder := new(strings.Builder)

	endpoints := route.EndpointInfo()
	maxMethodLength := maxMethodLength(endpoints)
	for _, endpoint := range endpoints {
		builder.WriteString(fmt.Sprintf("%-"+strconv.Itoa(maxMethodLength)+"s %s", endpoint.Method, endpoint.Path))

		if len(endpoint.Queries) > 0 {
			queryBuilder := new(strings.Builder)
			var queries []string
			for _, query := range endpoint.Queries {
				queryBuilder.WriteString(query.Name)
				if !query.Required {
					queryBuilder.WriteString("?")
				}

				queryBuilder.WriteString("=")
				if query.Value != "" {
					queryBuilder.WriteString(query.Value)
				}

				queries = append(queries, queryBuilder.String())
				builder.WriteString(fmt.Sprintf("?%s", strings.Join(queries, "&")))
			}
		}

		builder.WriteString("\n")
	}

	return builder.String()
}

func maxMethodLength(endpoints []ave.RouteInfo) int {
	max := 0
	for _, endpoint := range endpoints {
		if len(endpoint.Method) > max {
			max = len(endpoint.Method)
		}
	}

	return max
}

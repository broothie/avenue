package ave

import (
	"net/http"
	"path"

	"github.com/broothie/avenue/openapi"
	"github.com/go-openapi/runtime/middleware"
)

func (r *Route) Doc() {
	opts := middleware.RedocOpts{BasePath: r.path, SpecURL: path.Join(r.path, "/specs/{_:.*}")}

	r.
		Method(http.MethodGet).
		Path("/specs/{_:.*}").
		Handler(openapi.SpecHandler(r.Root()))

	r.
		Method(http.MethodGet).
		Handler(middleware.Redoc(opts, http.NotFoundHandler()))
}

func (r *Route) DocOmit() *Route {
	route := r.newChild()
	route.documentation.Skip = true
	return route
}

func (r *Route) DocSummary(summary string) *Route {
	route := r.newChild()
	route.documentation.Summary = summary
	return route
}

func (r *Route) DocDescription(description string) *Route {
	route := r.newChild()
	route.documentation.Description = description
	return route
}

func (r *Route) DocBody(keys ...Key) *Route {
	route := r.newChild()
	route.documentation.Body = append(route.documentation.Body, keys...)
	return route
}

func (r *Route) DocResponses(responses ...Response) *Route {
	route := r.newChild()
	route.documentation.Responses = append(route.documentation.Responses, responses...)
	return route
}

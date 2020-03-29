package ave

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

package ave

type RouteInfo struct {
	Method        string
	Path          string
	Queries       []Query
	Headers       []Header
	Documentation Documentation
}

func (r *Route) Info() RouteInfo {
	return RouteInfo{
		Method:        r.method,
		Path:          r.path,
		Queries:       r.queries,
		Headers:       r.headers,
		Documentation: r.documentation,
	}
}

func (r *Route) EndpointInfo() []RouteInfo {
	exportedRoutes := make([]RouteInfo, len(r.endpoints))
	for i, endpoint := range r.endpoints {
		exportedRoutes[i] = endpoint.Info()
	}

	return exportedRoutes
}

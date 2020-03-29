package endpoint

import ave "github.com/broothie/avenue"

type EndpointInfoer interface {
	EndpointInfo() []ave.RouteInfo
}

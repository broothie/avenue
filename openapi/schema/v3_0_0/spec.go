package v3_0_0

import "github.com/broothie/drroute/openapi/schema"

type Spec struct {
	OpenAPI string      `json:"openapi" yaml:"openapi"`
	Info    schema.Info `json:"info" yaml:"info"`
}

type Endpoint struct {
}

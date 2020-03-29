package v2_0

import "github.com/broothie/drroute/openapi/schema"

type Spec struct {
	Swagger string      `json:"swagger" yaml:"swagger"`
	Info    schema.Info `json:"info" yaml:"info"`
}

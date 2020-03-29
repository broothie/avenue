package v3_0_0

import "github.com/broothie/avenue/openapi/schema"

type (
	Spec struct {
		OpenAPI string      `json:"openapi" yaml:"openapi"`
		Info    schema.Info `json:"info" yaml:"info"`
		Paths   Paths       `json:"paths" yaml:"paths"`
	}

	Paths     map[string]Endpoints
	Endpoints map[string]Endpoint

	Endpoint struct {
		Summary     string              `json:"summary" yaml:"summary"`
		Description string              `json:"description" yaml:"description"`
		Parameters  []Parameter         `json:"parameters,omitempty" yaml:"parameters,omitempty"`
		Responses   map[string]Response `json:"responses,omitempty" yaml:"responses,omitempty"`
	}

	Parameter struct {
		In          string             `json:"in" yaml:"in"`
		Name        string             `json:"name" yaml:"name"`
		Description string             `json:"description" yaml:"description"`
		Schema      ParameterSchema    `json:"schema,omitempty" yaml:"schema,omitempty"`
		Content     map[string]Content `json:"content,omitempty" yaml:"content,omitempty"`
		Required    bool               `json:"required" yaml:"required"`
	}

	ParameterSchema struct {
		Type    string `json:"type" yaml:"type"`
		Minimum int    `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	}

	Content struct {
	}
)

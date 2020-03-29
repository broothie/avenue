package v2_0

import "github.com/broothie/avenue/openapi/schema"

type (
	Spec struct {
		Swagger string      `json:"swagger" yaml:"swagger"`
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
		In               string        `json:"in" yaml:"in"`
		Name             string        `json:"name" yaml:"name"`
		Description      string        `json:"description" yaml:"description"`
		Type             string        `json:"type,omitempty" yaml:"type,omitempty"`
		Required         bool          `json:"required" yaml:"required"`
		Minimum          *int          `json:"minimum,omitempty" yaml:"minimum,omitempty"`
		Default          interface{}   `json:"default,omitempty" yaml:"default,omitempty"`
		Enum             []interface{} `json:"enum,omitempty" yaml:"enum,omitempty"`
		CollectionFormat string        `json:"collectionFormat,omitempty" yaml:"collectionFormat,omitempty"`
		Items            Items         `json:"items,omitempty" yaml:"items,omitempty"`
		MinItems         *int          `json:"minItems,omitempty" yaml:"minItems,omitempty"`
		MaxItems         *int          `json:"maxItems,omitempty" yaml:"maxItems,omitempty"`
		UniqueItems      *bool         `json:"uniqueItems,omitempty" yaml:"uniqueItems,omitempty"`
		AllowEmptyValue  *bool         `json:"allowEmptyValue,omitempty" yaml:"allowEmptyValue,omitempty"`
	}

	Items struct {
		Type string        `json:"type" yaml:"type"`
		Enum []interface{} `json:"enum,omitempty" yaml:"enum,omitempty"`
	}

	Response struct {
		Description string         `json:"description" yaml:"description"`
		Schema      ResponseSchema `json:"schema" yaml:"schema"`
	}

	ResponseSchema struct {
		Type       string              `json:"type" yaml:"type"`
		Required   []string            `json:"required,omitempty" yaml:"required,omitempty"`
		Properties map[string]Property `json:"properties" yaml:"properties"`
	}

	Property struct {
		Type        string `json:"type" yaml:"type"`
		Description string `json:"description" yaml:"description"`
	}
)

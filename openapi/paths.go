package openapi

import (
	"strconv"
	"strings"

	ave "github.com/broothie/avenue"
)

type (
	Paths     map[string]Endpoints
	Endpoints map[string]Endpoint

	Endpoint struct {
		Summary     string              `json:"summary" yaml:"summary"`
		Description string              `json:"description" yaml:"description"`
		Parameters  []Parameter         `json:"parameters,omitempty" yaml:"parameters,omitempty"`
		Responses   map[string]Response `json:"responses,omitempty" yaml:"responses,omitempty"`
	}

	Parameter struct {
		In       string `json:"in" yaml:"in"`
		Name     string `json:"name" yaml:"name"`
		Type     string `json:"type,omitempty" yaml:"type,omitempty"`
		Schema   Schema `json:"schema,omitempty" yaml:"schema,omitempty"`
		Required bool   `json:"required" yaml:"required"`
	}

	Response struct {
		Description string            `json:"description,omitempty" yaml:"description,omitempty"`
		Content     map[string]Schema `json:"content,omitempty" yaml:"content,omitempty"`
	}

	Schema struct {
		Type    string `json:"type" yaml:"type"`
		Example string `json:"example,omitempty" yaml:"example,omitempty"`
	}
)

func makePaths(route *ave.Route, version Version) Paths {
	endpoints := route.EndpointInfo()

	paths := make(Paths)
	for _, endpoint := range endpoints {
		if endpoint.Documentation.Skip {
			continue
		}

		endpoints, ok := paths[endpoint.Path]
		if !ok {
			endpoints = make(Endpoints)
			paths[endpoint.Path] = endpoints
		}

		var parameters []Parameter
		for _, segment := range strings.Split(endpoint.Path, "/") {
			if len(segment) > 0 && segment[0] == '{' && segment[len(segment)-1] == '}' {
				name := strings.Split(strings.Trim(segment, "{}"), ":")[0]
				parameters = append(parameters, Parameter{In: "path", Name: name, Required: true})
			}
		}

		for _, query := range endpoint.Queries {
			param := Parameter{In: "query", Name: query.Name, Required: query.Required}
			parameters = append(parameters, param.WithType(query.Type, version))
		}

		for _, header := range endpoint.Headers {
			param := Parameter{In: "header", Name: header.Name, Required: header.Required}
			parameters = append(parameters, param.WithType(header.Type, version))
		}

		for _, key := range endpoint.Documentation.Body {
			param := Parameter{In: "body", Name: key.Name, Required: key.Required}
			parameters = append(parameters, param.WithType(key.Type, version))
		}

		responses := make(map[string]Response)
		for _, response := range endpoint.Documentation.Responses {
			contents := make(map[string]Schema)
			for key, schema := range response.Content {
				contents[key] = Schema{
					Type:    schema.Type,
					Example: schema.Example,
				}
			}

			responses[strconv.Itoa(response.Status)] = Response{
				Description: response.Description,
				Content:     contents,
			}
		}

		endpoints[strings.ToLower(endpoint.Method)] = Endpoint{
			Summary:     endpoint.Documentation.Summary,
			Description: endpoint.Documentation.Description,
			Parameters:  parameters,
			Responses:   responses,
		}
	}

	return paths
}

func (p Parameter) WithType(typ string, version Version) Parameter {
	switch version {
	case V2_0:
		p.Type = typ
	case V3_0_0:
		p.Schema.Type = typ
	}

	return p
}

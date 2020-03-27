package drr

import (
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type (
	PathRoot struct {
		Paths Paths `yaml:"paths"`
	}

	Paths     map[string]Endpoints
	Endpoints map[string]Endpoint

	Endpoint struct {
		Summary     string      `yaml:"summary"`
		Description string      `yaml:"description"`
		Parameters  []Parameter `yaml:"parameters,omitempty"`
	}

	Parameter struct {
		In       string `yaml:"in"`
		Name     string `yaml:"name"`
		Required bool   `yaml:"required"`
	}
)

func (r *Route) GenerateDoc() error {
	paths := make(Paths)
	for _, route := range r.routes {
		endpoints, ok := paths[route.path]
		if !ok {
			endpoints = make(Endpoints)
			paths[route.path] = endpoints
		}

		var parameters []Parameter
		for _, segment := range strings.Split(route.path, "/") {
			if len(segment) > 0 && segment[0] == '{' && segment[len(segment)-1] == '}' {
				name := strings.Split(strings.Trim(segment, "{}"), ":")[0]
				parameters = append(parameters, Parameter{In: "path", Name: name, Required: true})
			}
		}

		for _, query := range route.queries {
			parameters = append(parameters, Parameter{In: "query", Name: query.Key, Required: query.Required})
		}

		for _, header := range route.headers {
			parameters = append(parameters, Parameter{In: "header", Name: header.Key, Required: header.Required})
		}

		endpoints[strings.ToLower(route.method)] = Endpoint{
			Summary:     route.summary,
			Description: route.description,
			Parameters:  parameters,
		}
	}

	openapi, err := yaml.Marshal(PathRoot{Paths: paths})
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile("openapi.yml", openapi, os.ModePerm); err != nil {
		return err
	}

	return nil
}

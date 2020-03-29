package openapi

import (
	"io/ioutil"
	"os"
	"sort"

	ave "github.com/broothie/avenue"
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
)

func Generate(route *ave.Route, options ...Option) error {
	var opts Options
	for _, option := range options {
		opts = option(opts)
	}

	return generateFile(makePaths(route, opts.version()), opts)
}

func generateFile(paths Paths, options Options) error {
	filepath := options.fileAndPath()
	format := options.format()
	version := options.version()

	var dst Spec
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		dst = freshSpec(version)
	} else {
		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			return err
		}

		unmarshaler := format.Unmarshaler()
		if err := unmarshaler(data, &dst); err != nil {
			return err
		}
	}

	if err := mergo.Merge(&dst, Spec{Paths: paths}, mergo.WithOverride); err != nil {
		return errors.Wrap(err, "")
	}

	for path, dstEndpoints := range dst.Paths {
		srcEndpoints, ok := paths[path]
		if !ok {
			continue
		}

		for method, dstEndpoint := range dstEndpoints {
			srcEndpoint, ok := srcEndpoints[method]
			if !ok {
				continue
			}

			parameters, err := mergeParams(dstEndpoint.Parameters, srcEndpoint.Parameters)
			if err != nil {
				return err
			}

			dstEndpoint.Parameters = parameters
		}
	}

	marshaler := format.Marshaler()
	data, err := marshaler(dst)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filepath, data, os.ModePerm)
}

func mergeParams(dst, src []Parameter) ([]Parameter, error) {
	namedParams := make(map[string]*Parameter)
	for _, param := range src {
		namedParams[param.Name] = &param
	}

	for _, param := range dst {
		dstParam, exists := namedParams[param.Name]
		if !exists {
			namedParams[param.Name] = &param
			continue
		}

		if err := mergo.Merge(dstParam, param); err != nil {
			return nil, err
		}
	}

	params := make([]Parameter, len(namedParams))
	counter := 0
	for _, param := range namedParams {
		params[counter] = *param
		counter++
	}

	sort.SliceStable(params, func(i, j int) bool {
		return indexOfParameterByName(params[i].Name, src) < indexOfParameterByName(params[j].Name, src)
	})

	return params, nil
}

func indexOfParameterByName(name string, params []Parameter) int {
	for i, param := range params {
		if param.Name == name {
			return i
		}
	}

	return -1
}

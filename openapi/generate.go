package openapi

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/broothie/avenue/endpoint"
)

func GenerateFile(route endpoint.EndpointInfoer, options ...Option) error {
	var opts Options
	for _, option := range options {
		opts = option(opts)
	}

	return generateFile(makePaths(route, opts.version()), opts)
}

func SpecHandler(route endpoint.EndpointInfoer, options ...Option) http.HandlerFunc {
	var opts Options
	for _, option := range options {
		opts = option(opts)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		segments := strings.Split(r.URL.Path, "/")
		opts = OptionFilename(segments[len(segments)-1])(opts)

		data, err := generateFreshSpecData(makePaths(route, opts.version()), opts)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(data)
	}
}

func generateFile(paths Paths, options Options) error {
	filepath := options.fileAndPath()
	format := options.format()

	var data []byte
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		if data, err = generateFreshSpecData(paths, options); err != nil {
			return err
		}
	} else {
		fileData, err := ioutil.ReadFile(filepath)
		if err != nil {
			return err
		}

		var spec Spec
		unmarshaler := format.Unmarshaler()
		if err := unmarshaler(fileData, &spec); err != nil {
			return err
		}

		marshaler := format.Marshaler()
		if data, err = marshaler(spec); err != nil {
			return err
		}
	}

	return ioutil.WriteFile(filepath, data, os.ModePerm)
}

func generateFreshSpecData(paths Paths, options Options) ([]byte, error) {
	spec := freshSpec(options.version())
	spec.Paths = paths
	marshaler := options.format().Marshaler()
	data, err := marshaler(spec)
	if err != nil {
		return nil, err
	}

	return data, nil
}

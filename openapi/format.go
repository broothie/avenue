package openapi

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

type Format string

const (
	FormatInvalid Format = "INVALID"
	FormatJSON    Format = "json"
	FormatYAML    Format = "yml"
)

func (f Format) Unmarshaler() func([]byte, interface{}) error {
	switch f {
	case FormatJSON:
		return json.Unmarshal
	case FormatYAML:
		return yaml.Unmarshal
	default:
		return nil
	}
}

func (f Format) Marshaler() func(interface{}) ([]byte, error) {
	switch f {
	case FormatJSON:
		return jsonMarshalIndent
	case FormatYAML:
		return yaml.Marshal
	default:
		return nil
	}
}

func (f Format) Ext() string {
	switch f {
	case FormatJSON:
		return ".json"
	default:
		return ".yml"
	}
}

func jsonMarshalIndent(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

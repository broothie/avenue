package openapi

import (
	"path"
	"strings"
)

type Filename string

func (f Filename) String() string {
	return string(f)
}

func (f Filename) Base() string {
	return path.Base(f.String())
}

func (f Filename) Ext() string {
	return path.Ext(f.String())
}

func (f Filename) Dir() string {
	return path.Dir(f.String())
}

func (f Filename) Format() Format {
	switch f.Ext() {
	case ".json":
		return FormatJSON
	case ".yaml", ".yml":
		return FormatYAML
	default:
		return FormatInvalid
	}
}

func (f Filename) Version() Version {
	nonExtFilename := strings.TrimSuffix(f.Base(), f.Ext())
	switch nonExtFilename {
	case "swagger":
		return V2_0
	case "openapi":
		return V3_0_0
	default:
		return VersionInvalid
	}
}

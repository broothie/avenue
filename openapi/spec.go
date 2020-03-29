package openapi

type (
	Spec struct {
		OpenAPI Version `json:"openapi,omitempty" yaml:"openapi,omitempty"`
		Swagger Version `json:"swagger,omitempty" yaml:"swagger,omitempty"`
		Info    Info    `json:"info" yaml:"info"`
		Paths   Paths   `json:"paths" yaml:"paths"`
	}

	Info struct {
		Title       string `json:"title" yaml:"title"`
		Description string `json:"description" yaml:"description"`
		Version     string `json:"version" yaml:"version"`
	}
)

func freshSpec(version Version) Spec {
	switch version {
	case V2_0:
		return Spec{
			Swagger: V2_0,
			Info:    info(),
		}

	default:
		return Spec{
			Swagger: V3_0_0,
			Info:    info(),
		}
	}
}

func info() Info {
	return Info{
		Title:       "",
		Description: "",
		Version:     "0.0.1",
	}
}

package openapi

type Version string

const (
	VersionInvalid Version = "INVALID"
	V2_0           Version = "2.0"
	V3_0_0         Version = "3.0.0"
)

func (v Version) FileBasename() string {
	switch v {
	case V2_0:
		return "swagger"
	default:
		return "openapi"
	}
}

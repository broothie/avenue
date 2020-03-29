package ave

type (
	Documentation struct {
		Skip        bool
		Summary     string
		Description string
		Body        []Key
		Responses   []Response
	}

	Response struct {
		Status      int
		Description string
		Content     map[string]Schema
	}

	Schema struct {
		Type    string
		Example string
	}

	Keys []Key

	Key struct {
		Name     string
		Type     string
		Required bool
	}
)

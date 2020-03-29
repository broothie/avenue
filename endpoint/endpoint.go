package endpoint

type Endpoint struct {
	Method        string
	Path          string
	Queries       []Query
	Headers       []Header
	Documentation Documentation
}

type (
	Documentation struct {
		Skip        bool
		Summary     string
		Description string
		Body        []Key
		Responses   []Response
	}

	Key struct {
		Name     string
		Type     string
		Required bool
	}

	Query  Pair
	Header Pair
	Pair   struct {
		Name     string
		Value    string
		Type     string
		Required bool
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
)

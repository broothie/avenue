package openapi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParameter_WithType(t *testing.T) {
	t.Run(string(V2_0), func(t *testing.T) {
		param1 := Parameter{Name: "name"}
		param2 := param1.WithType("type 2", V2_0)

		assert.Equal(t, param1.Name, "name")
		assert.Equal(t, param2.Name, "name")

		assert.Equal(t, param1.Type, "")
		assert.Equal(t, param2.Type, "type 2")
	})

	t.Run(string(V3_0_0), func(t *testing.T) {
		param1 := Parameter{Name: "name"}
		param2 := param1.WithType("type 2", V3_0_0)

		assert.Equal(t, param1.Name, "name")
		assert.Equal(t, param2.Name, "name")

		assert.Equal(t, param1.Schema.Type, "")
		assert.Equal(t, param2.Schema.Type, "type 2")
	})
}

package json

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJson(t *testing.T) {
	data := make(map[string]interface{}, 0)
	data["key"] = 2.8265426280397245e17

	got, err := Marshal(data)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "{\"key\":2.8265426280397245e+17}", string(got))
}

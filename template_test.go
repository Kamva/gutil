package gutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderTextTemplate(t *testing.T) {
	var testTable = []struct {
		tag      string
		text     string
		data     interface{}
		result   string
		hasError bool
	}{
		{"t1", "", nil, "", false},
		{"t2", "hi {{.Name}}", map[string]string{"Name": "Mehran"}, "hi Mehran", false},
		{"t3", "hi {{.LastName}}", map[string]string{"Name": "mehran"}, "hi <no value>", false},
	}
	for _, data := range testTable {
		result, err := RenderText(data.text, data.data)
		if data.hasError {
			assert.Error(t, err, data.tag)
			continue
		}
		assert.Nil(t, err)
		assert.Equal(t, data.result, result, data.tag)
	}
}

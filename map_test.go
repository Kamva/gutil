package gutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKeyValuesToMap(t *testing.T) {
	list := []interface{}{"a", "b", "1", "2"}
	m := map[string]interface{}{
		"a": "b",
		"1": "2",
	}

	result, err := KeyValuesToMap(list...)

	if assert.Nil(t, err) {
		assert.Equal(t, m, result)
	}

}

func TestMapToKeyValue(t *testing.T) {
	m := map[string]interface{}{
		"1": "2",
		"a": "b",
	}
	list := []interface{}{ "1", "2","a", "b"}

	result := MapToKeyValue(m)

	assert.Equal(t, list, result)

}

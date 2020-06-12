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
	list := []interface{}{"1", "2", "a", "b"}

	result := MapToKeyValue(m)

	assert.Equal(t, list, result)
}

func TestMapKeysExtractor_Extract(t *testing.T) {
	m := map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"d": 1,
			},
			"c": 1,
		},
	}

	depths := []struct {
		tag   string
		depth int
		keys  []string
	}{
		{"t1", 1, []string{"a"}},
		{"t2", 2, []string{"a", "a.b", "a.c"}},
		{"t3", 3, []string{"a", "a.b", "a.b.d", "a.c"}},
		{"t4", 10, []string{"a", "a.b", "a.b.d", "a.c"}},
	}

	for _, inp := range depths {
		keys := MapPathExtractor{Depth: inp.depth, Separator: "."}.Extract(m)
		assert.Equal(t, inp.keys, keys, inp.tag)
	}

}

func TestMapKeysExtractor_ExtractInvalidDepth(t *testing.T) {
	m := map[string]interface{}{
		"a": map[string]interface{}{
			"b": map[string]interface{}{
				"d": 1,
			},
			"c": 1,
		},
	}

	keys := MapPathExtractor{Depth: 0, Separator: "."}.Extract(m)
	assert.Nil(t, keys)

}

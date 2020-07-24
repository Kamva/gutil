package gutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniqueStrings(t *testing.T) {
	var testTable = []struct {
		tag     string
		strings []string
		unique  []string
	}{
		{"t1", nil, nil},
		{"t2", []string{}, []string{}},
		{"t3", []string{"a", "b"}, []string{"a", "b"}},
		{"t4", []string{"a", "b", "b", "c"}, []string{"a", "b", "c"}},
	}
	for _, data := range testTable {
		assert.Equal(t, data.unique, UniqueStrings(data.strings), data.tag)
	}
}

func TestRemoveFromStrings(t *testing.T) {
	var testTable = []struct {
		tag       string
		strings   []string
		removal   []string
		cleanList []string
	}{
		{"t1", nil, nil, nil},
		{"t2", nil, []string{}, nil},
		{"t3", nil, []string{"a"}, nil},
		{"t4", []string{}, []string{}, []string{}},
		{"t5", []string{"a"}, nil, []string{"a"}},
		{"t6", []string{"a", "b"}, []string{"b"}, []string{"a"}},
		{"t7", []string{"a", "b"}, []string{"b", "b"}, []string{"a"}},
		{"t8", []string{"a", "b", "b", "c"}, []string{"a", "b", "c"}, []string{}},
	}
	for _, data := range testTable {
		assert.Equal(t, data.cleanList, RemoveFromStrings(data.strings, data.removal...), data.tag)
	}
}

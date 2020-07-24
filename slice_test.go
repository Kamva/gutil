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

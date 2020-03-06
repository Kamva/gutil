package gutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Contains check that needle exists in the haystack or not.
func TestContains(t *testing.T) {
	var testTable = []struct {
		tag      string
		haystack []string
		needle   string
		contains bool
	}{
		{"t1", []string{}, "f", false},
		{"t2", []string{"a", "b", "c"}, "a", true},
		{"t3", []string{"a", "b", "c"}, "d", false},
		{"t44", []string{"c", "a", "b"}, "b", true},
	}
	for _, data := range testTable {
		assert.Equal(t, data.contains, Contains(data.haystack, data.needle), data.tag)
	}
}

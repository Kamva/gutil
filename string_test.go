package gutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnyString(t *testing.T) {
	assert.Equal(t, "", AnyString())
	assert.Equal(t, "", AnyString("", "", ""))
	assert.Equal(t, "a", AnyString("a", "b"))
	assert.Equal(t, "b", AnyString("b", "a"))
	assert.Equal(t, "a", AnyString("a", "b", "c"))
}

func TestSub(t *testing.T) {
	table := []struct {
		Tag  string
		S1   []string
		S2   []string
		Diff []string
	}{
		{"t1", nil, nil, []string{}},
		{"t2", []string{}, nil, []string{}},
		{"t3", nil, []string{}, []string{}},
		{"t4", []string{"a", "b"}, []string{}, []string{"a", "b"}},
		{"t5", []string{"a", "b"}, []string{"a"}, []string{"b"}},
		{"t6", []string{"a", "b"}, []string{"a", "c"}, []string{"b"}},
		{"t7", []string{"a", "b"}, []string{"a", "b"}, []string{}},
	}

	for _, v := range table {
		t.Run(v.Tag, func(t *testing.T) {
			assert.Equal(t, v.Diff, Sub(v.S1, v.S2))
		})
	}
}

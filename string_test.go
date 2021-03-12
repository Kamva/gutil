package gutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AnyStringTest(t *testing.T) {
	assert.Equal(t, "", AnyString())
	assert.Equal(t, "", AnyString("", "", ""))
	assert.Equal(t, "a", AnyString("a", "b"))
	assert.Equal(t, "b", AnyString("b", "a"))
	assert.Equal(t, "a", AnyString("a", "b", "c"))
}

package gutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndirectValue(t *testing.T) {
	val := "abc"
	indirectVal, isNil := IndirectValue(&val)
	assert.Equal(t, indirectVal, val)
	assert.False(t, isNil)
}

func TestIndirectValueWithoutPointer(t *testing.T) {
	val := "abc"
	indirectVal, isNil := IndirectValue(val)
	assert.Equal(t, indirectVal, val)
	assert.False(t, isNil)
}

func TestIndirectValueNil(t *testing.T) {
	var val *string
	indirectVal, isNil := IndirectValue(val)
	assert.Equal(t, indirectVal, nil)
	assert.True(t, isNil)
}

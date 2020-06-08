package gutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndirectValue(t *testing.T) {
	val := "abc"
	indirectVal := IndirectValue(&val)
	assert.Equal(t, indirectVal, val)
}

func TestIndirectValueWithoutPointer(t *testing.T) {
	val := "abc"
	indirectVal := IndirectValue(val)
	assert.Equal(t, indirectVal, val)
}

func TestIndirectValueNil(t *testing.T) {
	var val *string
	indirectVal := IndirectValue(val)
	assert.Nil(t, indirectVal)
}

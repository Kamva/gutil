package gutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandStringLength(t *testing.T) {
	str := RandString(4)
	assert.Equal(t, len(str), 4)
}

func TestRandStringUniqueness(t *testing.T) {
	values := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		str := RandString(10)
		assert.NotContains(t, values, str)
		values[i] = str
	}
}

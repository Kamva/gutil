package gutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	F0 Bits = 1 << iota
	F1
	F2
)

func TestBits_Set(t *testing.T) {
	var b Bits
	b = b.Set(F0)
	b = b.Set(F1)
	assert.Equal(t, Bits(3), b)
}

func TestBits_Has(t *testing.T) {
	var b Bits = 6 // F1,F2
	assert.False(t, b.Has(F0))
	assert.True(t, b.Has(F1))
	assert.True(t, b.Has(F2))
}

func TestBits_Toggle(t *testing.T) {
	var b Bits = 5 // F0, F2
	b = b.Toggle(F0)
	b = b.Toggle(F1)
	assert.Equal(t, Bits(6), b)
}

func TestBits_Clear(t *testing.T) {
	var b Bits = 5 // F0, F2
	b = b.Clear(F0)
	assert.Equal(t, Bits(4), b)
}

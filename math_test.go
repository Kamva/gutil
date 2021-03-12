package gutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPercentOf(t *testing.T) {
	assert.Equal(t, float64(0), PercentOf(0, 50))
	assert.Equal(t, float64(5), PercentOf(100, 5))
	assert.Equal(t, float64(50), PercentOf(100, 50))
}

func TestPercent(t *testing.T) {
	assert.Equal(t, float64(0), Percent(0, 20))
	assert.Equal(t, float64(20), Percent(100, 20))
	assert.Equal(t, float64(5), Percent(100, 5))
}

func TestPercentInt(t *testing.T) {
	assert.Equal(t, float64(0), PercentInt(0, 20))
	assert.Equal(t, float64(20), PercentInt(100, 20))
	assert.Equal(t, float64(40), PercentInt(200, 20))
}

func TestPercentFloat(t *testing.T) {
	assert.Equal(t, float64(0), PercentFloat(0, 20))
	assert.Equal(t, float64(20), PercentFloat(100, 20))
	assert.Equal(t, float64(40), PercentFloat(200, 20))
}

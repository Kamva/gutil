package gutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTime(t *testing.T) {
	now := time.Now()
	assert.Equal(t, NewTime(now), &now)
}

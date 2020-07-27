package gutil

import (
	crand "crypto/rand"
	"encoding/binary"
	"errors"
	"math/rand"
)

var rnd = rand.New(cryptoSource{})

// SecureRandNumber generate secure random number in range of [min,max).
func SecureRandNumber(min, max int) int {
	if min > max {
		PanicErr(errors.New("min can not be greater than max"))
	}
	return rnd.Intn(max-min) + min
}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	_ = binary.Read(crand.Reader, binary.BigEndian, &v)
	return v
}

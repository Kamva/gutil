package gutil

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"math/big"
)

type RandType int

const (
	RandTypeAlphaNum RandType = iota + 1
	RandTypeAlpha
	RandTypeNumber
)

// RandNumber generate secure random number in range of [min,max).
func RandNumber(min, max int64) int64 {
	if min > max {
		PanicErr(errors.New("min can not be greater than max"))
	}
	nBig, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		PanicErr(err)
	}
	return nBig.Int64() + min
}

// RandString returns random string with length of n.
func RandString(n int) string {
	rb := make([]byte, n)
	_, err := rand.Read(rb)

	if err != nil {
		panic(err)
	}

	return base64.URLEncoding.EncodeToString(rb)[:n]
}

func RandStringWithType(size int, t RandType) string {
	var dictionary string
	switch t {
	case RandTypeAlphaNum:
		dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	case RandTypeAlpha:
		dictionary = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	case RandTypeNumber:
		dictionary = "0123456789"
	}

	var bytes = make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

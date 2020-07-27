package gutil

import (
	"math/rand"
	"time"
)

var randSeeded = false
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func init() {
	seedRand()
}

func seedRand() {
	rand.Seed(time.Now().UTC().UnixNano())
	randSeeded = true
}

// RandString returns random string with length of n.
func RandString(n int) string {
	if !randSeeded {
		seedRand()
	}
	return randSeq(n)
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

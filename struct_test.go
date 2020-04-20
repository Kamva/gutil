package gutil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestUnmarshalStruct(t *testing.T) {
	u := User{
		Name: "abc",
		Age:  2,
	}
	p := Person{}

	result := Person{
		Name: "abc",
		Age:  2,
	}

	if assert.Nil(t, UnmarshalStruct(u, &p)) {
		assert.Equal(t, result, p)
	}
}

package gutil

import (
	"errors"
	"reflect"
)

// IsNil function check value is nil or no. To check real value of interface
// is nil or not, should using reflection, check this
// https://play.golang.org/p/Isoo0CcAvr. Firstly check
// `val==nil` because reflection can not get value of
// zero val.
func IsNil(val interface{}) (result bool) {

	if val == nil {
		return true
	}

	switch v := reflect.ValueOf(val); v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.UnsafePointer,
		reflect.Interface, reflect.Slice:
		return v.IsNil()
	}

	return
}

// IndirectType returns indirect value's reflection type.
func IndirectType(val interface{}) (reflect.Type, error) {
	if IsNil(val) {
		return nil, errors.New("value can not be nil")
	}

	rType := reflect.TypeOf(val)

	if rType.Kind() == reflect.Ptr {
		rType = rType.Elem()
	}

	return rType, nil
}

// NewInstanceByValue get the value and return new instance of that value type.
// Provided value can be either by reference or by value.
func NewInstanceByValue(v interface{}) (interface{}, error) {
	t, err := IndirectType(v)
	if err != nil {
		return nil, err
	}

	return reflect.New(t).Interface(), nil
}

// StructTags return struct all fields tags.
func StructTags(val interface{}) ([]reflect.StructTag, error) {
	rType, err := IndirectType(val)

	if err != nil {
		return nil, err
	}

	if rType.Kind() != reflect.Struct {
		return nil, errors.New("value must be a struct or pointer to struct")
	}

	tags := make([]reflect.StructTag, rType.NumField())
	for i := 0; i < rType.NumField(); i++ {
		tags = append(tags, rType.Field(i).Tag)
	}

	return tags, nil
}

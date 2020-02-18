package gutil

// Must check if error is not nil, so panic, otherwise return value.
func Must(val interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}

	return val
}

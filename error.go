package gutil

func Must(val interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}

	return val
}

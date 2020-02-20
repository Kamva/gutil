package gutil

// PanicIfNil will panic error param if value is nil.
func PanicIfNil(val interface{}, err error) {
	if IsNil(val) {
		panic(err)
	}
}

package gutil

// Must check if error is not nil, so panic, otherwise return value.
func Must(val interface{}, err error) interface{} {
	if err != nil {
		panic(err)
	}

	return val
}

// AnyErr returns first error that is not nil.
func AnyErr(errors ...error) error {
	for _, err := range errors {
		if !IsNil(err) {
			return err
		}
	}

	return nil
}

func CauseErr(err error) error {
	for {
		u, ok := err.(interface {
			Unwrap() error
		})
		if !ok {
			break
		}
		err = u.Unwrap()
	}
	return err
}

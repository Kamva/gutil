package gutil

import (
	"errors"
	"fmt"
)

func KeyValuesToMap(keyValues ...interface{}) (map[string]interface{}, error) {
	m := make(map[string]interface{})

	for i := 0; i < len(keyValues); i += 2 {
		if i+1 < len(keyValues) {
			m[fmt.Sprint(keyValues[i])] = keyValues[i+1]
		} else {
			return nil, errors.New(fmt.Sprintf("missing %v index value", i))
		}
	}

	return m, nil
}

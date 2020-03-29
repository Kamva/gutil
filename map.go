package gutil

import (
	"errors"
	"fmt"
)

// KeyValuesToMap convert each pair of slice to a key/value of map.
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

// MapToKeyValue convert map to slice that contain map key and value pairs.
func MapToKeyValue(m map[string]interface{}) []interface{} {
	kv := make([]interface{}, len(m)*2)

	i := 0
	for k, v := range m {
		kv[i] = k
		kv[i+1] = v
		i += 2
	}

	return kv
}

// ExtendMap extend dest map by src map.
// overwrite argument set overwrite policy.
func ExtendMap(src, dest map[string]interface{}, overwrite bool) {
	for key, val := range src {
		// If key exists in dest and we can not overwrite it, so continue.
		if _, ok := dest[key]; ok && !overwrite {
			continue
		}
		dest[key] = val
	}
}

// ExtendMap extend dest map by src map.
// overwrite argument set overwrite policy.
func ExtendStrMap(src, dest map[string]string, overwrite bool) {
	for key, val := range src {
		// If key exists in dest and we can not overwrite it, so continue.
		if _, ok := dest[key]; ok && !overwrite {
			continue
		}
		dest[key] = val
	}
}

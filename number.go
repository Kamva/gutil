package gutil

import "strconv"

// ParseInt parse integer and returns value if string
// value is a valid integer or default value
func ParseInt(val string, def int) int {
	if val == "" {
		return def
	}
	if result, err := strconv.Atoi(val); err == nil {
		return result
	}
	return def
}

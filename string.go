package gutil

// StringDefault return default value if provided value is empty.
func StringDefault(val, def string) string {
	if val == "" {
		return def
	}

	return val
}

package gutil

// BoolPtr returns pointer to the bool value.
func BoolPtr(b bool) *bool {
	return &b
}

// IntPtr returns pointer to the integer value.
func IntPtr(i int) *int {
	return &i
}

// Int64Ptr returns pointer to the int64 value.
func Int64Ptr(i int64) *int64 {
	return &i
}

// StringPtr returns pointer to the string value.
func StringPtr(s string) *string {
	return &s
}

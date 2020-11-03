package gutil

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

// StringDefault return default value if provided value is empty.
func StringDefault(val, def string) string {
	if val == "" {
		return def
	}

	return val
}

// ToSnakeCase returns snake_case of the provided value.
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// ReplaceAt replace provided string in specific index to another index.
// removal part is [begin,end).
// e.g., to replace char "a" in the "salam" word, call ReplaceAt("salam","b",1,2) // result will be "sblam".
func ReplaceAt(str string, replace string, begin int, end int) string {
	return str[:begin] + replace + str[end:]
}

// ReplaceRune replace single rune in specific index
func ReplaceRune(str string, r rune, removalIndex int) string {
	return ReplaceAt(str, string(r), removalIndex, removalIndex+1)
}

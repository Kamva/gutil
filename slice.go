package gutil

func UniqueStrings(values []string) []string {
	if values==nil {
		return nil
	}

	unique := make([]string, 0)
	for _, v := range values {
		if !Contains(unique, v) {
			unique = append(unique, v)
		}
	}

	return unique
}

package util

func FilterSlice[T any](s []T, test func(T) bool) []T {
	var filtered []T
	for _, v := range s {
		if test(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

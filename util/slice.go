package util

import "math/rand"

func FilterSlice[T any](s []T, test func(T) bool) []T {
	var filtered []T
	for _, v := range s {
		if test(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func SampleSlice[T any](s []T) T {
	idx := rand.Intn(len(s))
	return s[idx]
}

func Sample[T any](s ...T) T {
	return SampleSlice(s)
}

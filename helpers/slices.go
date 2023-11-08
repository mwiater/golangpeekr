// Package common provides general utility functions for slice operations.
package helpers

// SliceContains checks if an item is present in the given slice.
func SliceContains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

// SliceIntersection returns a new slice containing the common elements of two slices.
func SliceIntersection[T comparable](slice1, slice2 []T) []T {
	set := make(map[T]struct{})
	var intersection []T

	for _, v := range slice1 {
		set[v] = struct{}{}
	}

	for _, v := range slice2 {
		if _, found := set[v]; found {
			intersection = append(intersection, v)
			delete(set, v) // Optional: to prevent duplicates if slice2 contains duplicates
		}
	}

	return intersection
}

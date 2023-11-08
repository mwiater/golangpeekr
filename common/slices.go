// Package common provides general utility functions for slice operations.
package common

import (
	"fmt"
	"reflect"
)

// UniqueSlice takes a slice and returns a new slice with unique elements.
func UniqueSlice[T comparable](slice []T) []T {
	uniqMap := make(map[T]struct{})
	uniqSlice := []T{}

	for _, item := range slice {
		if _, exists := uniqMap[item]; !exists {
			uniqMap[item] = struct{}{}
			uniqSlice = append(uniqSlice, item)
		}
	}

	return uniqSlice
}

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

// SliceFindFieldValue searches for structs in a slice that have a given field with a specified value.
// It returns a slice of structs that match the criteria.
func SliceFindFieldValue(slice interface{}, fieldName string, value interface{}) ([]interface{}, error) {
	// Ensure input is a slice
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return nil, fmt.Errorf("expected a slice, got %s", rv.Kind())
	}

	var result []interface{}

	// Iterate over the slice
	for i := 0; i < rv.Len(); i++ {
		item := rv.Index(i)
		// Ensure the item is a struct
		if item.Kind() != reflect.Struct {
			continue
		}

		// Get the field by name
		field := item.FieldByName(fieldName)
		if !field.IsValid() {
			continue // Field not found, skip
		}

		// Compare the field value
		if reflect.DeepEqual(field.Interface(), value) {
			result = append(result, item.Interface())
		}
	}

	return result, nil
}

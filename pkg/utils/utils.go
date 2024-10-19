package utils

import (
	"errors"
	"reflect"
)

// Function that expects a pointer type
func ExpectPointer(v interface{}) error {
	// Use reflection to check if v is a pointer
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return errors.New("expected a pointer")
	}

	return nil
}

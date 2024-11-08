package utils

import (
	"errors"
	"reflect"
)

// Function that expects a pointer type
func ExpectPointer(vs ...interface{}) error {
	// Use reflection to check if v is a pointer
	for _, v := range vs {
		if reflect.ValueOf(v).Kind() != reflect.Ptr {
			return errors.New("expected a pointer")
		}
	}
	return nil
}

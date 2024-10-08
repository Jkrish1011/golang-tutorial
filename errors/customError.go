package errors

import (
	"errors"
	"fmt"
)

type userError struct {
	name string
}

// Just implement the Error() function so indirectly it implements the Error interface
// Just one function is required: Error()
func (e userError) Error() string {
	return fmt.Sprintf("%v has a problem with their account", e.name)
}

// Even this is enough to create a custom error type

var err error = errors.New("something went wrong")

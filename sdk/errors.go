package sdk

import (
	"errors"
	"fmt"
)

// Define a sentinel error for "website not found"
var WebsiteNotFoundError = errors.New("3as website not found")

// Custom error type to add context around the WebsiteNotFoundError
type NotFoundError struct {
	Resource string
	Err      error
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("resource %s: %v", e.Resource, e.Err)
}

func (e *NotFoundError) Unwrap() error {
	return e.Err
}

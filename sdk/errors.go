package sdk

import (
	"errors"
	"fmt"
)

var NotFoundErr = errors.New("resource not found")

type NotFoundError struct {
	Resource string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s: %s", e.Resource, NotFoundErr.Error())
}

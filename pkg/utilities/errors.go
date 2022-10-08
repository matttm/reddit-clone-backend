package errors

import (
	"fmt"
)

type BadRequestError struct {
	ErrorString string
}

func (e *BadRequestError) Error() string {
	return fmt.Sprintf("Error: %s", e.ErrorString)
}

type GenericError struct {
	ErrorString string
}

func (e *GenericError) Error() string {
	return fmt.Sprintf("Error: %s", e.ErrorString)
}

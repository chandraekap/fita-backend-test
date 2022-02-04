package errors

import "fmt"

type NotFoundError struct {
	Model string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s: was not found.", e.Model)
}

func NewNotFoundError(model string) *NotFoundError {
	return &NotFoundError{
		Model: model,
	}
}

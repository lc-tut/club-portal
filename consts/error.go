package consts

import "errors"

type ValidateError struct {
	text string
}

func (err *ValidateError) Error() string {
	return err.text
}

func NewValidateError(text string) *ValidateError {
	return &ValidateError{text: text}
}

var UnreachableError = errors.New("unreachable")

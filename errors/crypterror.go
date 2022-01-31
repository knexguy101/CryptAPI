package errors

import "errors"

type ErrorData struct {
	CryptError CryptError
	InnerError error
}

type CryptError error

// NewErrorData creates a new ErrorData object. Should only be used to create new static returns.
func NewErrorData(err CryptError, innerErr error) *ErrorData {
	return &ErrorData{
		CryptError: err,
		InnerError: innerErr,
	}
}

func NewCryptError(msg string) CryptError {
	return errors.New(msg)
}
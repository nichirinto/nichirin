package exception

import "errors"

type Exception struct {
	Code    string
	Message string
	Err     error
}

func NewWithErr(code string, message string, err error) *Exception {
	return &Exception{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func New(code string, message string) *Exception {
	e := errors.New(message)

	return &Exception{
		Code:    code,
		Message: message,
		Err:     e,
	}
}

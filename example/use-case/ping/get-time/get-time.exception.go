package gettime

import "github.com/nichirinto/nichirin/framework/common/exception"

const InvalidInput = "INVALID_INPUT"

func NewInvalidInputException() exception.Exception {
	return exception.NewWithCode(InvalidInput)
}

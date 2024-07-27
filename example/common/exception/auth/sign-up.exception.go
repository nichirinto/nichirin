package authexception

import "github.com/nichirinto/nichirin/framework/common/exception"

const UsernameExistedErrorCode = "USERNAME_EXISTED"

func NewUsernameExistedException() exception.Exception {
	return exception.NewWithCode(UsernameExistedErrorCode)
}

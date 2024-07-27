package signupusecase

import (
	authexception "github.com/nichirinto/nichirin/example/common/exception/auth"
	"github.com/nichirinto/nichirin/example/storage/db/repo"
	"github.com/nichirinto/nichirin/framework/common/exception"
)

func checkUsernameExisted(c context) exception.Exception {
	existed, e := repo.User.FindByEmail(c.Input.Username)

	if e != nil {
		return exception.ServerError
	}
	if existed != nil {
		return authexception.NewUsernameExistedException()
	}

	return nil
}

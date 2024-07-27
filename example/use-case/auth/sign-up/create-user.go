package signupusecase

import (
	"github.com/nichirinto/nichirin/example/storage/db/repo"
	userrepo "github.com/nichirinto/nichirin/example/storage/db/repo/user"
	"github.com/nichirinto/nichirin/example/storage/db/schema"
	"github.com/nichirinto/nichirin/framework/common/exception"
)

func createUser(c context) (*schema.User, exception.Exception) {
	user, e := repo.User.CreateNew(&userrepo.CreateUserInput{
		Username: c.Input.Username,
		Password: c.Input.Password,
	})

	return user, e
}

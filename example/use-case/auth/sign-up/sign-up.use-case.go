package signupusecase

import (
	authdto "github.com/nichirinto/nichirin/example/common/dto/auth"
	authexception "github.com/nichirinto/nichirin/example/common/exception/auth"
	"github.com/nichirinto/nichirin/example/storage/db/repo"
	userrepo "github.com/nichirinto/nichirin/example/storage/db/repo/user"
	"github.com/nichirinto/nichirin/framework/common/exception"
	"github.com/nichirinto/nichirin/framework/core"
)

func Handle(c *core.Context[authdto.SignUpInputDto]) *core.Res[authdto.SignUpOutputDto] {
	c.Logger.Info(c.Input.Password)

	res := &authdto.SignUpOutputDto{}

	existed, e := repo.User.FindByEmail(c.Input.Username)

	if e != nil {
		return core.CreateResponse(res, exception.ServerError)
	}
	if existed != nil {
		return core.CreateResponse(res, authexception.NewUsernameExistedException())
	}

	user, e := repo.User.CreateNew(&userrepo.CreateUserInput{
		Username: c.Input.Username,
		Password: c.Input.Password,
	})
	if e != nil {
		return core.CreateResponse(res, authexception.NewUsernameExistedException())
	}

	c.Logger.Info(user.Id)

	res.UserId = user.Id

	return core.CreateResponse(res, nil)
}

package signupusecase

import (
	authdto "github.com/nichirinto/nichirin/example/common/dto/auth"
	"github.com/nichirinto/nichirin/framework/core"
)

func Handle(c context) *core.Res[authdto.SignUpOutputDto] {
	c.Logger.Info(c.Input.Password)

	res := &authdto.SignUpOutputDto{}

	e := checkUsernameExisted(c)
	if e != nil {
		return core.CreateResponse(res, e)
	}

	user, e := createUser(c)
	if e != nil {
		return core.CreateResponse(res, e)
	}

	res.UserId = user.Id

	return core.CreateResponse(res, nil)
}

package controller

import (
	authdto "github.com/nichirinto/nichirin/example/common/dto/auth"
	signupusecase "github.com/nichirinto/nichirin/example/use-case/auth/sign-up"
	"github.com/nichirinto/nichirin/framework/core"
)

func Auth(app *core.Nichirin) {
	new(core.Router[authdto.SignUpInputDto, authdto.SignUpOutputDto]).Path("/auth/sign-up").Handle(signupusecase.Handle).Attach(app)
}

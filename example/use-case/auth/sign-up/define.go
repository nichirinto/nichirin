package signupusecase

import (
	authdto "github.com/nichirinto/nichirin/example/common/dto/auth"
	"github.com/nichirinto/nichirin/framework/core"
)

type context = *core.Context[authdto.SignUpInputDto]

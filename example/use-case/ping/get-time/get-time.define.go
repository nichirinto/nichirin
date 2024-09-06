package gettime

import "github.com/nichirinto/nichirin/framework/core"

type InputDto struct {
	Unix string `json:"unix" validate:"required,min=1"`
}

type OutputDto struct {
	Time string `json:"time"`
}

type Ctx = *core.Context[InputDto, OutputDto]

type IInput struct {
	unix int64
}

type IOutput = *core.Res[OutputDto]

type IRouter = core.Router[InputDto, OutputDto]

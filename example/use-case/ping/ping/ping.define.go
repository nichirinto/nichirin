package ping

import "github.com/nichirinto/nichirin/framework/core"

type Request struct{}

type Response struct {
	Message string `json:"message"`
}

type Ctx = *core.Context[Request, Response]

type Input struct{}

type Output = *core.Res[Response]

type Router = core.Router[Request, Response]

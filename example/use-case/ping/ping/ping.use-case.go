package ping

import (
	"github.com/nichirinto/nichirin/example/engine"
	"github.com/nichirinto/nichirin/framework/core"
)

func Handle() {
	new(Router).Register(engine.App, core.Get, "/", handle)
}

func handle(ctx Ctx) Output {
	return ctx.Return(&Response{Message: "pong"})
}

package gettime

import (
	"github.com/nichirinto/nichirin/example/engine"
	"github.com/nichirinto/nichirin/framework/common/exception"
	"github.com/nichirinto/nichirin/framework/core"
	"github.com/nichirinto/nichirin/framework/lib/datetime"
	"strconv"
)

func Handle() {
	new(IRouter).Register(engine.App, core.Get, "/time", handle)
}

func validateInput(ctx Ctx) (*IInput, exception.Exception) {
	result := new(IInput)

	unix, err := strconv.ParseInt(ctx.Input.Unix, 10, 64)

	if err != nil {
		return nil, NewInvalidInputException()
	}

	result.unix = unix

	return result, nil
}

func handle(ctx Ctx) IOutput {
	input, err := validateInput(ctx)
	if err != nil {
		return ctx.Throw(err)
	}

	ctx.Output.Time = datetime.UnixToFormat(input.unix)

	return ctx.Return()
}

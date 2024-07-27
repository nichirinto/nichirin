package controller

import (
	pingdto "github.com/nichirinto/nichirin/example/common/dto/ping"
	pingusecase "github.com/nichirinto/nichirin/example/use-case/ping/ping"
	"github.com/nichirinto/nichirin/framework/core"
)

func Ping(app *core.Nichirin) {
	new(core.Router[pingdto.PingInputDto, pingdto.PingOutputDto]).Path("/").Handle(pingusecase.Handle).Attach(app)
}

package ping

import (
	"github.com/nichirinto/nichirin/framework/common/dto"
	"github.com/nichirinto/nichirin/framework/core"
)

type PingOutputDto struct {
	Pong bool `json:"pong"`
}

func PingRoute(app *core.Nichirin, url string) {
	new(core.Router[dto.EmptyDto, PingOutputDto]).Path(url).Handle(func(*core.Context[dto.EmptyDto]) *core.Res[PingOutputDto] {
		return core.CreateResponse(&PingOutputDto{
			Pong: true,
		}, nil)
	}).Attach(app)
}

package ping

import "github.com/nichirinto/nichirin/framework/core"

type AuthInputDto struct{}

type AuthOutputDto struct {
	Pong bool `json:"pong"`
}

func PingRoute(app *core.Nichirin, url string) {
	new(core.Router[AuthInputDto, AuthOutputDto]).Path(url).Handle(func(*core.Context[AuthInputDto]) *core.Res[AuthOutputDto] {
		return core.CreateResponse(&AuthOutputDto{
			Pong: true,
		}, nil)
	}).Attach(app)
}

package ping

import "github.com/nichirinto/nichirin/framework/core"

type AuthInputDto struct{}

type AuthOutputDto struct {
	Pong bool `json:"pong"`
}

func AuthUseCase(c *core.Context[AuthInputDto]) *core.Res[AuthOutputDto] {
	return core.CreateResponse(&AuthOutputDto{
		Pong: true,
	}, nil)
}

func AuthRoute(app *core.Nichirin, url string) {
	new(core.Router[AuthInputDto, AuthOutputDto]).Path(url).Handle(AuthUseCase).Attach(app)
}

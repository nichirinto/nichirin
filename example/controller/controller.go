package controller

import "github.com/nichirinto/nichirin/framework/core"

func Init(app *core.Nichirin) {
	Auth(app)
	Ping(app)
}

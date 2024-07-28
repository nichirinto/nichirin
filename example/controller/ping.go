package controller

import (
	"github.com/nichirinto/nichirin/framework/core"
	"github.com/nichirinto/nichirin/framework/handler/ping"
)

func Ping(app *core.Nichirin) {
	ping.PingRoute(app, "/")
}

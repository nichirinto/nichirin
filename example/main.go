package main

import (
	"github.com/nichirinto/nichirin/example/controller"
	_ "github.com/nichirinto/nichirin/example/controller"
	"github.com/nichirinto/nichirin/example/storage/db"
	"github.com/nichirinto/nichirin/framework/core"
)

func main() {
	db.Init()

	app := core.NewApp("127.0.0.1:8888")

	controller.Init(app)

	app.StartServer()
}

package main

import (
	_ "github.com/nichirinto/nichirin/example/controller"
	"github.com/nichirinto/nichirin/example/engine"
)

func main() {
	engine.App.StartServer()
}

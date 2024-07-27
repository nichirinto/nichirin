package db

import (
	"github.com/nichirinto/nichirin/example/storage/db/repo"
	"github.com/nichirinto/nichirin/framework/storage/pg"
)

var Db *pg.NichirinPg

func Init() {
	db, err := pg.NewPg("postgres://admin:123456@localhost:5432/test")

	if err != nil {
		panic(err)
	}

	Db = db

	repo.Init(Db)

	Db.Logger.Info("DB connected")
}

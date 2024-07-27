package repo

import (
	userrepo "github.com/nichirinto/nichirin/example/storage/db/repo/user"
	"github.com/nichirinto/nichirin/framework/storage/pg"
)

var User userrepo.UserRepo

func Init(db *pg.NichirinPg) {
	User = userrepo.Init(db)
}

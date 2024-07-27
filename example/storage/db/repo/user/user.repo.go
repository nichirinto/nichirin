package userrepo

import (
	"github.com/nichirinto/nichirin/example/storage/db/schema"
	"github.com/nichirinto/nichirin/framework/common/exception"
	"github.com/nichirinto/nichirin/framework/lib/uid"
	"github.com/nichirinto/nichirin/framework/storage/pg"
)

type UserRepo struct {
	pg.NichirinRepo
}

func Init(db *pg.NichirinPg) UserRepo {
	repo := UserRepo{}

	repo.Pg = db
	repo.Db = db.Db

	return repo
}

func (r UserRepo) FindByEmail(username string) (*schema.User, exception.Exception) {
	var user *schema.User

	result := r.Db.First(&user, "username = ?", username)

	if pg.RecordNotFound(result) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, exception.DbException
	}

	return user, nil
}

func (r UserRepo) CreateNew(input *CreateUserInput) (*schema.User, exception.Exception) {
	newUser := &schema.User{Id: uid.New(), Username: input.Username, Password: input.Password}

	result := r.Db.Create(newUser)

	if result.Error != nil {
		return nil, exception.NewDbException(result.Error)
	}

	return newUser, nil
}

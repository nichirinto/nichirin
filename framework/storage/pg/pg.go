package pg

import (
	"github.com/nichirinto/nichirin/framework/common/exception"
	"github.com/nichirinto/nichirin/framework/lib/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (r *NichirinPg) Printf(msg string, data ...interface{}) {
	r.Logger.Info(msg)
}

func NewPg(url string) (*NichirinPg, exception.Exception) {
	pg := &NichirinPg{
		Url: url,
	}

	pg.Logger = logger.Log.With()

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil || db == nil {
		return nil, ThrowDbError(err)
	}

	sqlDb, _ := db.DB()

	pg.Sql = sqlDb
	pg.Db = db

	return pg, nil
}

func (db *NichirinPg) Ping() exception.Exception {
	if err := db.Sql.Ping(); err != nil {
		return ThrowDbError(err)
	}

	return nil
}

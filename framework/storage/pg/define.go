package pg

import (
	"database/sql"
	"github.com/nichirinto/nichirin/framework/lib/logger"
	"gorm.io/gorm"
)

type Db = *gorm.DB
type Sql = *sql.DB

type NichirinPg struct {
	Url    string
	Sql    *sql.DB
	Db     *gorm.DB
	Logger *logger.Logger
}

type NichirinRepo struct {
	Pg *NichirinPg
	Db *gorm.DB
}

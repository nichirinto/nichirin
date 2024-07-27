package pg

import (
	"github.com/nichirinto/nichirin/framework/common/exception"
	"net/http"
)

func ThrowDbError(err error) exception.Exception {
	e := exception.Init()
	e.StatusCode = http.StatusInternalServerError
	e.Code = exception.PgxError
	e.Err = err
	e.Message = err.Error()

	return e
}

func RecordNotFound(tx Db) bool {
	return tx.Error != nil && tx.Error.Error() == "record not found"
}

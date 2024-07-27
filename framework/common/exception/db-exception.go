package exception

const RecordNotFoundErrorCode = "RECORD_NOT_FOUND"

var RecordNotFoundException = NewWithCode(RecordNotFoundErrorCode)

const DbExceptionErrorCode = "DB_ERROR"

var DbException = NewWithCode(DbExceptionErrorCode)

func NewDbException(err error) Exception {
	return NewWithErr(DbExceptionErrorCode, DbExceptionErrorCode, err)
}

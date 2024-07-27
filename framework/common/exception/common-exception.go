package exception

const EngineErrorCode = "ENGINE_ERROR"
const ServerErrorCode = "SERVER_ERROR"
const BadRequestErrorCode = "BAD_REQUEST"
const PgxError = "PGX_ERROR"

var EngineError = New(EngineErrorCode, "Sql Error")
var ServerError = New(ServerErrorCode, "Bad Request")

package exception

const EngineErrorCode = "ENGINE_ERROR"
const ServerErrorCode = "SERVER_ERROR"

var EngineError = New(EngineErrorCode, "Engine Error")
var ServerError = New(ServerErrorCode, "Bad Request")

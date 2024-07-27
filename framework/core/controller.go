package core

import (
	"encoding/json"
	"github.com/nichirinto/nichirin/framework/common/exception"
	"net/http"
)

func ThrowBadRequest(w http.ResponseWriter, err error) {
	res := new(Res[interface{}])

	res.Err = exception.NewWithErr(exception.BadRequestErrorCode, err.Error(), err)

	SendResponse[interface{}](w, res)
}

func SendResponse[T any](w http.ResponseWriter, v *Res[T]) {
	if v.Err != nil {
		statusCode := v.Err.StatusCode

		if statusCode < 200 || statusCode >= 600 {
			statusCode = http.StatusBadRequest
		}

		w.WriteHeader(statusCode)
	}

	res := new(ResPayload[T])
	res.Data = v.Data
	res.Code = "OK"

	if v.Err != nil {
		res.Code = v.Err.Code
		res.Message = v.Err.Message
	}

	_ = json.NewEncoder(w).Encode(res)
}

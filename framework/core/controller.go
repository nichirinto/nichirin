package core

import (
	"encoding/json"
	"github.com/nichirinto/nichirin/framework/common/exception"
	"net/http"
)

type WebResponse[T any] struct {
	Code    string
	Data    *T
	Message string
}

func sendResponse[T any](w http.ResponseWriter, data *T, err *exception.Exception) {
	response := WebResponse[T]{Data: data}

	if err != nil {
		response.Code = err.Code
		response.Message = err.Message
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

func handle[I any, O any](controller Controller[I, O]) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		input := new(I)
		err := json.NewDecoder(r.Body).Decode(input)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
		}
	}
}

func InitController[I any, O any](controller Controller[I, O]) {

}

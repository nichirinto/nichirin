package core

import (
	"github.com/go-chi/chi/v5"
	"github.com/nichirinto/nichirin/framework/common/exception"
	"net/http"
)

type Nichirin struct {
	Engine *chi.Mux
}

type HttpMethod string

const (
	Get     HttpMethod = http.MethodGet
	Post    HttpMethod = http.MethodPost
	Put     HttpMethod = http.MethodPut
	Patch   HttpMethod = http.MethodPatch
	Delete  HttpMethod = http.MethodDelete
	Head    HttpMethod = http.MethodHead
	Options HttpMethod = http.MethodOptions
)

type Controller[I any, O any] struct {
	Url     string
	Method  HttpMethod
	Handler func(Nichirin, I) (O, exception.Exception)
}

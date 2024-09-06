package core

import (
	"github.com/go-chi/chi/v5"
	"github.com/nichirinto/nichirin/framework/common/exception"
	"github.com/nichirinto/nichirin/framework/lib/logger"
	"net/http"
)

type Nichirin struct {
	Engine  *chi.Mux
	Address string
}

type Router[TI any, TO any] struct {
	Name      string
	Method    HttpMethod
	Url       string
	AuthGuard bool
	Handler   func(*Context[TI, TO]) *Res[TO]
}

type Context[TI any, TO any] struct {
	Request   *http.Request
	TraceId   string
	Input     *TI
	Output    *TO
	Exception exception.Exception
	Logger    *logger.Logger

	Timeline ContextTimeline
}

type ContextTimeline struct {
	Started  int64
	Finished int64
}

type Res[T any] struct {
	Data *T                  `json:"data"`
	Err  exception.Exception `json:"error"`
}

type ResPayload[T any] struct {
	Code    string `json:"code"`
	Data    *T     `json:"data"`
	Message string `json:"message"`
}

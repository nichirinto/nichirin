package exception

import (
	"errors"
	"github.com/nichirinto/nichirin/framework/lib/logger"
	"go.uber.org/zap"
	"net/http"
)

type RuntimeException struct {
	Code       string
	Message    string
	Err        error
	StatusCode int
}

type Exception = *RuntimeException

func (r Exception) Log(logger *logger.Logger) {
	logger.Error(r.Message, zap.String("code", r.Code), zap.Any("status", r.StatusCode))
}

func Init() Exception {
	return &RuntimeException{}
}

func NewWithErr(code string, message string, err error) *RuntimeException {
	return &RuntimeException{
		Code:       code,
		Message:    message,
		Err:        err,
		StatusCode: http.StatusBadRequest,
	}
}

func New(code string, message string) *RuntimeException {
	e := errors.New(message)

	return &RuntimeException{
		Code:       code,
		Message:    message,
		Err:        e,
		StatusCode: http.StatusBadRequest,
	}
}

func NewWithCode(code string) *RuntimeException {
	e := errors.New(code)

	return &RuntimeException{
		Code:       code,
		Message:    code,
		Err:        e,
		StatusCode: http.StatusBadRequest,
	}
}

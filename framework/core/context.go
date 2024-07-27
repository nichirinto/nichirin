package core

import (
	"github.com/nichirinto/nichirin/framework/lib/datetime"
	"github.com/nichirinto/nichirin/framework/lib/logger"
	"github.com/nichirinto/nichirin/framework/lib/uid"
	"go.uber.org/zap"
	"net/http"
)

func NewContext[T any](req *http.Request) *Context[T] {
	ctx := new(Context[T])

	traceId := uid.New()
	started := datetime.Unix()

	ctx.Request = req
	ctx.TraceId = traceId

	ctx.Timeline = ContextTimeline{
		Started:  started,
		Finished: 0,
	}

	ctx.Logger = logger.Log.With(
		zap.String("context", traceId),
	)

	return ctx
}

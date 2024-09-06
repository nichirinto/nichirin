package core

import (
	"github.com/nichirinto/nichirin/framework/common/exception"
	"github.com/nichirinto/nichirin/framework/lib/datetime"
	"github.com/nichirinto/nichirin/framework/lib/logger"
	"github.com/nichirinto/nichirin/framework/lib/uid"
	"go.uber.org/zap"
	"net/http"
)

func NewContext[TI any, TO any](req *http.Request) *Context[TI, TO] {
	ctx := new(Context[TI, TO])

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

func (r *Context[TI, TO]) Return(data ...*TO) *Res[TO] {
	if len(data) > 0 {
		r.Output = data[0]
	}

	res := new(Res[TO])

	res.Data = r.Output
	res.Err = r.Exception

	if res.Err != nil {
		res.Data = nil
	}

	return res
}

func (r *Context[TI, TO]) Throw(err exception.Exception) *Res[TO] {
	r.Exception = err

	res := new(Res[TO])

	res.Err = r.Exception

	return res
}

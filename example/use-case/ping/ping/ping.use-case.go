package pingusecase

import (
	pingdto "github.com/nichirinto/nichirin/example/common/dto/ping"
	"github.com/nichirinto/nichirin/framework/core"
	"github.com/nichirinto/nichirin/framework/lib/datetime"
)

func Handle(context *core.Context[pingdto.PingInputDto]) *core.Res[pingdto.PingOutputDto] {
	res := &pingdto.PingOutputDto{
		Time: datetime.Unix(),
	}

	return core.CreateResponse(res, nil)
}

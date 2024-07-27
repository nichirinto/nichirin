package core

import "github.com/nichirinto/nichirin/framework/common/exception"

func CreateResponse[T any](data *T, err exception.Exception) *Res[T] {
	res := new(Res[T])

	res.Data = data
	res.Err = err

	if err != nil {
		res.Data = nil
	}

	return res
}

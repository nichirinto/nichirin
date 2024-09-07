package result

import "github.com/nichirinto/nichirin/framework/common/exception"

type Resolver[T any] struct {
	Data T
	E    exception.Exception
}

func Return[T any](data T, e exception.Exception) Resolver[T] {
	return Resolver[T]{Data: data, E: e}
}

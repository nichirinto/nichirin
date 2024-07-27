package core

import "net/http"

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

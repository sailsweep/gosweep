package types

import (
	serviceCtx "entdemo/svc"
)

type HandleFunc func(ctx *serviceCtx.ServiceContext, Req Request) (Response, error)

type Request struct {
	Body  any
	Query any
}

type Response struct {
	Code int
	Msg  string
	Data any
}

const (
	EntSuccess = 200
	EntError   = 400
	EntInvalid = 100
	EntToken   = 101
)

var CodeInfo = map[int]string{
	EntSuccess: "success",
	EntToken:   "invalid params",
	EntInvalid: "invalid token",
	EntError:   "server error",
}

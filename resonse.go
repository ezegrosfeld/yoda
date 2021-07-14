package yoda

import "github.com/savsgio/atreugo/v11"

func NewResponse(status int, message string, ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(status)
	resp := Response{Message: message}
	return ctx.JSONResponse(resp, status)
}

func JSON(status int, obj interface{}, ctx *atreugo.RequestCtx) error {
	ctx.SetStatusCode(status)
	return ctx.JSONResponse(obj, status)
}

func Created(obj interface{}, ctx *atreugo.RequestCtx) error {
	return JSON(201, obj, ctx)
}

func OK(obj interface{}, ctx *atreugo.RequestCtx) error {
	return JSON(200, obj, ctx)
}

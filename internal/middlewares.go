package internal

import (
	"time"

	"github.com/savsgio/atreugo/v11"
)

// HandlerInfo struct
type HandlerInfo struct {
	method  []byte
	path    []byte
	tsStart int64
}

// BeforeMiddleware is part of the logging middleware
func BeforeMiddleware() atreugo.Middleware {
	return func(ctx *atreugo.RequestCtx) error {
		ctx.SetUserValue("info", HandlerInfo{
			method:  ctx.Method(),
			path:    ctx.Path(),
			tsStart: time.Now().UnixNano(),
		})
		return ctx.Next()
	}
}

// AfterMiddleware is part of the logging middleware
func AfterMiddleware() atreugo.Middleware {
	return func(ctx *atreugo.RequestCtx) error {
		LogResponse(ctx)
		return ctx.Next()
	}
}

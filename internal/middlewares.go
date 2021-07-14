package internal

import (
	"time"

	"github.com/savsgio/atreugo/v11"
)

const (
	uvKey       = "request.startTime"
	millisecond = float64(time.Millisecond)
)

type HandlerInfo struct {
	method  []byte
	path    []byte
	tsStart int64
}

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

func AfterMiddleware() atreugo.Middleware {
	return func(ctx *atreugo.RequestCtx) error {
		LogResponse(ctx)
		return ctx.Next()
	}
}

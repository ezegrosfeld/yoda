package yoda

import (
	"encoding/json"

	"github.com/savsgio/atreugo/v11"
)

func Bind(i interface{}, ctx *atreugo.RequestCtx) error {
	return json.Unmarshal(ctx.PostBody(), i)
}

func BindFromContext(key string, ctx *atreugo.RequestCtx) interface{} {
	return ctx.UserValue(key)
}

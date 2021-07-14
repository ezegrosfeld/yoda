package yoda

import "github.com/savsgio/atreugo/v11"

func (r *Router) Get(path string, handler Handler) {
	r.Parent.GET(path, func(rc *atreugo.RequestCtx) error {
		ctx := &Context{rc}
		return handler(ctx)
	})
}

func (r *Router) Post(path string, handler Handler) {
	r.Parent.POST(path, func(rc *atreugo.RequestCtx) error {
		ctx := &Context{rc}
		return handler(ctx)
	})
}

func (r *Router) Put(path string, handler Handler) {
	r.Parent.PUT(path, func(rc *atreugo.RequestCtx) error {
		ctx := &Context{rc}
		return handler(ctx)
	})
}

func (r *Router) Delete(path string, handler Handler) {
	r.Parent.DELETE(path, func(rc *atreugo.RequestCtx) error {
		ctx := &Context{rc}
		return handler(ctx)
	})
}

func (r *Router) Patch(path string, handler Handler) {
	r.Parent.PATCH(path, func(rc *atreugo.RequestCtx) error {
		ctx := &Context{rc}
		return handler(ctx)
	})
}

// Append middlewares to router
func (r *Router) Use(middlewares ...Handler) {
	for _, fn := range middlewares {
		r.Parent.UseBefore(func(rc *atreugo.RequestCtx) error {
			ctx := &Context{rc}
			return fn(ctx)
		})
	}
}

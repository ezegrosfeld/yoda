package yoda

import "github.com/savsgio/atreugo/v11"

// Get creates a listener of GET requests
func (r *Router) Get(path string, handler Handler) {
	r.Parent.GET(path, func(rc *atreugo.RequestCtx) error {
		ctx := &Context{rc}
		return handler(ctx)
	})
}

// Post creates a listener of POST requests
func (r *Router) Post(path string, handler Handler) {
	r.Parent.POST(path, func(rc *atreugo.RequestCtx) error {
		ctx := &Context{rc}
		return handler(ctx)
	})
}

// Put creates a listener of PUT requests
func (r *Router) Put(path string, handler Handler) {
	r.Parent.PUT(path, func(rc *atreugo.RequestCtx) error {
		ctx := &Context{rc}
		return handler(ctx)
	})
}

// PATCH creates a listener of PATCH requests
func (r *Router) Patch(path string, handler Handler) {
	r.Parent.PATCH(path, func(rc *atreugo.RequestCtx) error {
		ctx := &Context{rc}
		return handler(ctx)
	})
}

// Delete creates a listener of DELETE requests
func (r *Router) Delete(path string, handler Handler) {
	r.Parent.DELETE(path, func(rc *atreugo.RequestCtx) error {
		ctx := &Context{rc}
		return handler(ctx)
	})
}

// Use appends middlewares to router
func (r *Router) Use(middlewares ...Handler) {
	for _, fn := range middlewares {
		r.Parent.UseBefore(func(rc *atreugo.RequestCtx) error {
			ctx := &Context{rc}
			return fn(ctx)
		})
	}
}

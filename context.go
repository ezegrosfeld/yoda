package yoda

import "context"

// Next function
func (c *Context) Next() error {
	return c.Parent.Next()
}

// GetHeader returns an specific value of a given key
func (c *Context) GetHeader(key string) []byte {
	return c.Parent.Request.Header.Peek(key)
}

// Param returns the value of a given param
func (c *Context) Param(key string) interface{} {
	return c.Parent.UserValue(key)
}

// Query returns the value of a given query key
func (c *Context) Query(key string) []byte {
	return c.Parent.QueryArgs().Peek(key)
}

// SetValue sets a value with a given key
func (c *Context) SetValue(key string, value interface{}) {
	c.Parent.SetUserValue(key, value)
}

// GetValue gets the value set by SetValue
func (c *Context) GetValue(key string) interface{} {
	return c.Parent.UserValue(key)
}

// AttachContext adds a new context to the current one
func (c *Context) AttachContext(context context.Context) {
	c.Parent.AttachContext(context)
}

// AttachedContext returns the attached context
func (c *Context) AttachedContext() context.Context {
	return c.Parent.AttachedContext()
}

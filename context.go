package yoda

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

package yoda

// Next function
func (c *Context) Next() error {
	return c.Parent.Next()
}

// Get a header value
func (c *Context) GetHeader(key string) []byte {
	return c.Parent.Request.Header.Peek(key)
}

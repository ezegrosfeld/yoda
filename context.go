package yoda

// Next function
func (c *Context) Next() error {
	return c.Parent.Next()
}

func (c *Context) GetHeader(key string) []byte {
	return c.Parent.Request.Header.Peek(key)
}

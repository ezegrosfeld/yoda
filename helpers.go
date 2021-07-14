package yoda

import (
	"encoding/json"
)

// Binds a json into a struct
func (c *Context) Bind(i interface{}) error {
	return json.Unmarshal(c.Parent.PostBody(), i)
}

// Returns a value from a context using the key
func (c *Context) FromContext(key string) interface{} {
	return c.Parent.UserValue(key)
}

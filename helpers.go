package yoda

import (
	"encoding/json"

	"github.com/go-playground/validator"
)

// Binds a json into a struct
func (c *Context) Bind(i interface{}) error {
	json.Unmarshal(c.Parent.PostBody(), i)
	validate := validator.New()
	return validate.Struct(i)
}

// Binds a json with custom validaroe
func (c *Context) BindWithValidator(i interface{}, key string, fn validator.Func) error {
	json.Unmarshal(c.Parent.PostBody(), i)
	validate := validator.New()
	validate.RegisterValidation(key, fn)
	return validate.Struct(i)
}

// Returns a value from a context using the key
func (c *Context) FromContext(key string) interface{} {
	return c.Parent.UserValue(key)
}

package yoda

import (
	"github.com/ezegrosfeld/vader"
)

// Creates a default response with a message
func (c *Context) NewResponse(status int, message string) error {
	c.Parent.SetStatusCode(status)
	resp := Response{Message: message}
	return c.Parent.JSONResponse(resp, status)
}

// Creates a JSON response
func (c *Context) JSON(status int, obj interface{}) error {
	c.Parent.SetStatusCode(status)
	return c.Parent.JSONResponse(obj, status)
}

// Creates a JSON response with status created
func (c *Context) Created(obj interface{}) error {
	return c.JSON(201, obj)
}

// Creates a JSON response with status OK
func (c *Context) OK(obj interface{}) error {
	return c.JSON(200, obj)
}

// Creates a JSON response with an error
func (c *Context) Error(err vader.Error) error {
	status := err.GetCode()
	resp := Response{Error: err.Error()}
	return c.JSON(status, resp)
}

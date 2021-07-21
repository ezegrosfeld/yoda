package yoda

import (
	"github.com/ezegrosfeld/vader"
)

// Message creates a default response with a message
func (c *Context) Message(status int, message string) error {
	c.Parent.SetStatusCode(status)
	resp := Response{Message: message}
	return c.Parent.JSONResponse(resp, status)
}

// JSON creates a JSON response
func (c *Context) JSON(status int, obj interface{}) error {
	resp := Response{Data: obj}
	c.Parent.SetStatusCode(status)
	return c.Parent.JSONResponse(resp, status)
}

// Error creates a JSON response with an error
func (c *Context) Error(err vader.Error) error {
	status := err.GetCode()
	resp := Response{Error: err.Error()}
	return c.Parent.JSONResponse(resp, status)
}

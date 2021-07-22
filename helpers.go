package yoda

import (
	"encoding/json"
	"fmt"

	"github.com/ezegrosfeld/vader"
	"github.com/go-playground/validator"
)

func handleValidatorError(errors validator.ValidationErrors) string {
	text := ""
	for _, err := range errors {
		if text == "" {
			text = fmt.Sprintf("Error in field %s of type %s", err.Field(), err.Kind())
		} else {
			text = fmt.Sprintf("%s, ", text)
		}
	}
	return text
}

// Bind binds a json into a struct
func (c *Context) Bind(i interface{}) vader.Error {
	jsonerr := json.Unmarshal(c.Parent.PostBody(), i)
	if jsonerr != nil {
		return vader.BadRequest(jsonerr.Error())
	}
	validate := validator.New()
	if validate.Struct(i) == nil {
		return nil
	}
	err := validate.Struct(i).(validator.ValidationErrors)
	fmt.Println(handleValidatorError(err))
	return vader.BadRequest(handleValidatorError(err))
}

// BindWithValidator binds a json with custom validaroe
func (c *Context) BindWithValidator(i interface{}, key string, fn validator.Func) vader.Error {
	jsonerr := json.Unmarshal(c.Parent.PostBody(), i)
	if jsonerr != nil {
		return vader.BadRequest(jsonerr.Error())
	}
	validate := validator.New()
	verr := validate.RegisterValidation(key, fn)
	if verr != nil {
		return vader.InternalError(verr.Error())
	}
	if validate.Struct(i) == nil {
		return nil
	}
	err := validate.Struct(i).(validator.ValidationErrors)
	fmt.Println(handleValidatorError(err))
	return vader.BadRequest(handleValidatorError(err))
}

// FromContext returns a value from a context using the key
func (c *Context) FromContext(key string) interface{} {
	return c.Parent.UserValue(key)
}

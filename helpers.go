package yoda

import (
	"encoding/json"
	"fmt"

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

// Binds a json into a struct
func (c *Context) Bind(i interface{}) error {
	jsonerr := json.Unmarshal(c.Parent.PostBody(), i)
	if jsonerr != nil {
		return fmt.Errorf(jsonerr.Error())
	}
	validate := validator.New()
	err := validate.Struct(i).(validator.ValidationErrors)
	if err != nil {
		fmt.Println(handleValidatorError(err))
		return fmt.Errorf(handleValidatorError(err))
	}
	return nil
}

// Binds a json with custom validaroe
func (c *Context) BindWithValidator(i interface{}, key string, fn validator.Func) error {
	jsonerr := json.Unmarshal(c.Parent.PostBody(), i)
	if jsonerr != nil {
		return fmt.Errorf(jsonerr.Error())
	}
	validate := validator.New()
	validate.RegisterValidation(key, fn)
	err := validate.Struct(i).(validator.ValidationErrors)
	if err != nil {
		return fmt.Errorf(handleValidatorError(err))
	}
	return nil
}

// Returns a value from a context using the key
func (c *Context) FromContext(key string) interface{} {
	return c.Parent.UserValue(key)
}

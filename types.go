package yoda

import "github.com/savsgio/atreugo/v11"

type Yoda interface {
	Run()
}

type yoda struct {
	*atreugo.Atreugo
}

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

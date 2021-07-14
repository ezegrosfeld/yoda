package yoda

import (
	"github.com/ezegrosfeld/yoda/internal"
	"github.com/savsgio/atreugo/v11"
)

func New(name string) *yoda {
	server := atreugo.New(atreugo.Config{
		Name: name,
		Addr: ":8080",
	})

	server.UseBefore(internal.BeforeMiddleware())
	server.UseAfter(internal.AfterMiddleware())

	return &yoda{server}
}

func (y *yoda) Run() {
	if err := y.ListenAndServe(); err != nil {
		panic(err)
	}
}

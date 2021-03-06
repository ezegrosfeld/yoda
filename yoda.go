package yoda

import (
	"os"

	"github.com/atreugo/cors"
	"github.com/ezegrosfeld/yoda/internal"
	"github.com/savsgio/atreugo/v11"
)

// NewServer creates a new standard yoda server
func NewServer(config Config) *Yoda {
	server := atreugo.New(atreugo.Config{
		Name:         config.Name,
		Addr:         config.Addr,
		IdleTimeout:  config.IdleTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
		LogAllErrors: os.Getenv("ENVIRONMENT") != "production",
	})

	corsConfig := config.CORS
	corsMiddleware := cors.New(corsConfig)

	server.UseAfter(corsMiddleware)
	server.UseBefore(internal.BeforeMiddleware())
	server.UseAfter(internal.AfterMiddleware())

	ir := &Router{server.Router}

	return &Yoda{server, ir}
}

//NewServerWithConfig creates a yoda server with config
func NewServerWithConfig(config atreugo.Config) *Yoda {
	server := atreugo.New(config)

	server.UseBefore(internal.BeforeMiddleware())
	server.UseAfter(internal.AfterMiddleware())

	ir := &Router{server.Router}

	return &Yoda{server, ir}
}

// Start starts the yoda server
func (y *Yoda) Start() error {
	return y.ListenAndServe()
}

// Group creates a group
func (y *Yoda) Group(path string, fns ...Handler) *Router {
	r := y.NewGroupPath(path)
	router := &Router{r}
	router.Use(fns...)
	return router
}

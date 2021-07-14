package yoda

import (
	"github.com/atreugo/cors"
	"github.com/ezegrosfeld/yoda/internal"
	"github.com/savsgio/atreugo/v11"
)

// Creates a new standard yoda server
func NewServer(config Config) *yoda {
	server := atreugo.New(atreugo.Config{
		Name:         config.Name,
		Addr:         config.Addr,
		IdleTimeout:  config.IdleTimeout,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	})

	corsConfig := config.CORS
	corsMiddleware := cors.New(corsConfig)

	server.UseAfter(corsMiddleware)
	server.UseBefore(internal.BeforeMiddleware())
	server.UseAfter(internal.AfterMiddleware())

	ir := &Router{server.Router}

	return &yoda{server, ir}
}

//Creates a yoda server with config
func NewServerWithConfig(config atreugo.Config) *yoda {
	server := atreugo.New(config)

	server.UseBefore(internal.BeforeMiddleware())
	server.UseAfter(internal.AfterMiddleware())

	ir := &Router{server.Router}

	return &yoda{server, ir}
}

// Starts the yoda server
func (y *yoda) Start() {
	if err := y.ListenAndServe(); err != nil {
		panic(err)
	}
}

// Creates a group
func (y *yoda) Group(path string, fns ...Handler) *Router {
	r := y.NewGroupPath(path)
	router := &Router{r}
	router.Use(fns...)
	return router
}

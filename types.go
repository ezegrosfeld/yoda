package yoda

import (
	"time"

	"github.com/atreugo/cors"
	"github.com/savsgio/atreugo/v11"
)

// Yoda is the server interface
type Yoda interface {
	Run()
}

// Server struct
type yoda struct {
	*atreugo.Atreugo
	internalRouter *Router
}

// Response struct
type Response struct {
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// Context struct
type Context struct {
	Parent *atreugo.RequestCtx
}

// Router struct
type Router struct {
	Parent *atreugo.Router
}

// Handler struct
type Handler func(*Context) error

// Config struct
type Config struct {
	Addr         string
	Name         string
	CORS         cors.Config
	IdleTimeout  time.Duration
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

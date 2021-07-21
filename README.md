[![Go Report Card](https://goreportcard.com/badge/github.com/ezegrosfeld/yoda)](https://goreportcard.com/report/github.com/ezegrosfeld/yoda)
[![GoDev](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/ezegrosfeld/yoda)
[![GitHub release](https://img.shields.io/github/release/ezegrosfeld/yoda.svg)](https://github.com/ezegrosfeld/yoda/releases)

# Yoda web framework

built on top of [atreugo](https://github.com/savsgio/atreugo).

Example server on localhost:8080:

```go
srv := yoda.NewServer(yoda.Config{
	Addr:   ":8080",
	Name:   "server",
})

group := srv.Group("/group")

group.Get("/create", func (ctx *yoda.Context) error {
    return ctx.JSON(200, "")
})

srv.StartServer()
```

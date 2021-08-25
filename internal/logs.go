package internal

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/savsgio/atreugo/v11"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[97;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

// LogResponse creates a coloured log of the response
func LogResponse(ctx *atreugo.RequestCtx) {
	env := os.Getenv("ENVIRONMENT")
	if env == "production" {
		return
	}
	status := ctx.Response.StatusCode()
	color := getResponseColor(status)

	t := time.Now()
	date := fmt.Sprintf("%d/%d/%d - %d:%d:%d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())

	ip := ctx.LocalIP()

	info := ctx.UserValue("info").(HandlerInfo)

	methodColor := strings.Replace(getMethodColor(string(info.method)), "4", "3", 1)

	timeHandler := (time.Now().UnixNano() - info.tsStart) / (int64(time.Millisecond) / int64(time.Nanosecond))

	fmt.Printf("%s%s%s %s [%s %3d %s] => %s | %d ms | %s\n", methodColor, info.method, reset, info.path, color, status, reset, date, timeHandler, ip)
}

func getResponseColor(status int) string {
	if status < 300 {
		return green
	}

	if status > 500 {
		return red
	}

	return yellow
}

func getMethodColor(method string) string {
	switch method {
	case "GET":
		return cyan
	case "POST":
		return yellow
	case "PUT":
		return blue
	case "PATCH":
		return magenta
	case "DELETE":
		return red
	default:
		return white
	}
}

package main

import (
	"github.com/bartam1/mobilfox/shorter/internal/handlers/httphandler"
	"github.com/bartam1/mobilfox/shorter/pkg/logs/extlog"
	"github.com/bartam1/mobilfox/shorter/pkg/logs/httplog"
	echo "github.com/labstack/echo/v4"
)

func main() {
	extlog.Init()

	e := echo.New()

	e.Use(httplog.MiddlewareLogging)
	e.HTTPErrorHandler = httplog.ErrorHandler

	httphandler.RegisterHandlers(e, nil)
	e.Start("0.0.0.0:7777")
}

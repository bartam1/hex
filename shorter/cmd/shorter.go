package main

import (
	"github.com/bartam1/mobilfox/shorter/internal/handlers/httphandler"
	echo "github.com/labstack/echo/v4"
	//echomiddleware "github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	httphandler.RegisterHandlers(e, nil)
	e.Start("0.0.0.0:7777")
}

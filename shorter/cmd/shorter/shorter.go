package main

import (
	"context"
	"os"

	port "github.com/bartam1/mobilfox/shorter/internal/core/ports"
	"github.com/bartam1/mobilfox/shorter/internal/core/services/command"
	"github.com/bartam1/mobilfox/shorter/internal/core/services/query"
	"github.com/bartam1/mobilfox/shorter/internal/handlers/httphandler"
	"github.com/bartam1/mobilfox/shorter/internal/repositories/memrepo"
	"github.com/bartam1/mobilfox/shorter/pkg/httpserver"
	"github.com/bartam1/mobilfox/shorter/pkg/logs/extlog"
	"github.com/bartam1/mobilfox/shorter/pkg/logs/httplog"
	echo "github.com/labstack/echo/v4"
)

func main() {
	extlog.Init()

	ctx := context.Background()

	repo, _ := memrepo.New()

	s := port.Service{
		Queries: port.Queries{
			UrlsWidthHash: query.NewUrlsWidthHash(repo),
			Url:           query.NewUrl(repo),
		},
		Commands: port.Commands{
			MakeUrlHash: command.NewMakeUrlHash(repo),
			DeleteUrl:   command.NewDeleteUrl(repo),
		},
	}
	hndl := httphandler.New(s)

	e := echo.New()

	e.Use(httplog.MiddlewareLogging)
	e.HTTPErrorHandler = httplog.ErrorHandler

	httphandler.RegisterHandlers(e, hndl)

	idleConnsClosed := make(chan struct{})
	go httpserver.CatchInterrupt(ctx, idleConnsClosed, e.Server)

	e.Start("0.0.0.0:" + os.Getenv("PORT"))
}

package main

import (
	"context"
	"os"

	port "github.com/bartam1/mobilfox/shorter/internal/core/ports"
	"github.com/bartam1/mobilfox/shorter/internal/core/services/command"
	"github.com/bartam1/mobilfox/shorter/internal/core/services/query"
	"github.com/bartam1/mobilfox/shorter/internal/handlers/httphandler"
	"github.com/bartam1/mobilfox/shorter/internal/repositories/psqlrepo"
	"github.com/bartam1/mobilfox/shorter/pkg/httpserver"
	"github.com/bartam1/mobilfox/shorter/pkg/logs/extlog"
	"github.com/bartam1/mobilfox/shorter/pkg/logs/httplog"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	extlog.Init()

	ctx := context.Background()

	//repo, _ := memrepo.New()

	repo, err := psqlrepo.New(ctx, "postgres://postgres:almafa@psql:5432/db")

	if err != nil {
		logrus.Panicf("db error: ", err)
	}

	//Add interactions with prev created repo
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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))
	//Add logrus middleware
	e.Use(httplog.MiddlewareLogging)
	e.HTTPErrorHandler = httplog.ErrorHandler

	httphandler.RegisterHandlers(e, hndl)

	//Catch interruptions and shutdown gracefully
	idleConnsClosed := make(chan struct{})
	go httpserver.CatchInterrupt(ctx, idleConnsClosed, e.Server)
	//Listen address and port (for docker-compose)
	e.Start("shorter:" + os.Getenv("PORT"))
}

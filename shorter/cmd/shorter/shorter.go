package main

import (
	"context"
	"os"

	port "github.com/bartam1/mobilfox/shorter/internal/core/ports"
	service "github.com/bartam1/mobilfox/shorter/internal/core/services"
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

	//repo, err := memrepo.New()

	repo, err := psqlrepo.New(ctx, os.Getenv("DATABASE_PSQL_URL"))

	if err != nil {
		logrus.Panicf("db error: ", err)
	}

	//Add interactions with prev created repo
	s := service.Shorter{
		Queries: service.Queries{
			UrlsWidthHash: query.NewUrlsWidthHash(repo),
			Url:           query.NewUrl(repo),
		},
		Commands: service.Commands{
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

	port.RegisterHandlers(e, hndl)

	//Catch interruptions and shutdown gracefully
	idleConnsClosed := make(chan struct{})
	go httpserver.CatchInterrupt(ctx, idleConnsClosed, e.Server)
	//Listen address and port (for docker-compose)
	e.Start(os.Getenv("HOST") + ":" + os.Getenv("PORT"))
}

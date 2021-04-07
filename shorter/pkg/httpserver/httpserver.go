package httpserver

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

func CatchInterrupt(ctx context.Context, idleConnsClosed chan struct{}, srv *http.Server) {
	sigint := make(chan os.Signal, 1)

	// interrupt signal sent from terminal
	signal.Notify(sigint, os.Interrupt)
	// sigterm signal sent from kubernetes
	signal.Notify(sigint, syscall.SIGTERM)

	<-sigint

	logrus.Info("Interrupt received => Server is shutting down gracefully")

	// We received an interrupt signal, shut down.
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		logrus.Printf("HTTP server Shutdown interrupted: %v", err)
	}
	close(idleConnsClosed)

}

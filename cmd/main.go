package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/go-github/v68/github"
	"github.com/woxQAQ/frp-webconsole/pkg/gen/frpc"
	frpcSvc "github.com/woxQAQ/frp-webconsole/pkg/gen/http/frpc/server"
	"github.com/woxQAQ/frp-webconsole/pkg/log"
	"github.com/woxQAQ/frp-webconsole/pkg/services"
	"github.com/woxQAQ/frp-webconsole/pkg/stores"
	goahttp "goa.design/goa/v3/http"

	"go.uber.org/zap"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := log.Logger()
	logger.Info("Starting server")

	frpcService := services.NewFrpcService(stores.NewGithubClient(github.NewClient(nil)))
	endpoints := frpc.NewEndpoints(frpcService)
	mux := goahttp.NewMuxer()
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)
	frpcsvc := frpcSvc.New(endpoints, mux, dec, enc, nil, nil)
	frpcsvc.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		})
	})
	frpcSvc.Mount(mux, frpcsvc)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("Failed to start server", zap.Error(err))
		}
	}()

	<-ctx.Done()
	logger.Info("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Failed to shutdown server", zap.Error(err))
	}
	logger.Info("Server shutdown complete")
}

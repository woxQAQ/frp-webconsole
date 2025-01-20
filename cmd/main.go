package main

import (
	"context"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-fuego/fuego"

	"github.com/woxQAQ/frp-webconsole/pkg/controller"
	"github.com/woxQAQ/frp-webconsole/pkg/router"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	svc := fuego.NewServer(func(s *fuego.Server) {
		s.Addr = "localhost:8080"
	}, fuego.WithOpenAPIConfig(fuego.OpenAPIConfig{
		PrettyFormatJson: true,
		JsonFilePath:     "../api/openapi.json",
	}))
	r := router.NewRouter(svc)
	r.Register(controller.NewFrpController())
	go func() {
		if err := svc.Run(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server", slog.Any("error", err))
		}
	}()
	<-ctx.Done()
	slog.Info("Shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svc.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", slog.Any("error", err))
	}
	slog.Info("Server shutdown complete")
}

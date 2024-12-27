package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/woxQAQ/frp-webconsole/pkg/log"
	"github.com/woxQAQ/frp-webconsole/pkg/middleware"
	"go.uber.org/zap"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	logger := log.Logger()
	logger.Info("Starting server")

	engine := gin.New()
	engine.Use(middleware.Logger(logger), middleware.Recovery(logger))
	engine.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
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

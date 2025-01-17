package main

import (
	"context"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	_ "github.com/woxQAQ/frp-webconsole/api"
	"github.com/woxQAQ/frp-webconsole/pkg/controller"
	"github.com/woxQAQ/frp-webconsole/pkg/log"
	"github.com/woxQAQ/frp-webconsole/pkg/middleware"
	"github.com/woxQAQ/frp-webconsole/pkg/router"
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
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router := router.NewRouter(engine)
	router.Register(controller.NewFrpController())

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

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

	"github.com/woxQAQ/frp-webconsole/pkg/controller"
	"github.com/woxQAQ/frp-webconsole/pkg/log"
	"github.com/woxQAQ/frp-webconsole/pkg/middleware"
	"github.com/woxQAQ/frp-webconsole/pkg/router"
	"go.uber.org/zap"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used

//	@securitydefinitions.oauth2.application	OAuth2Application
//	@tokenUrl								https://example.com/oauth/token
//	@scope.write							Grants write access
//	@scope.admin							Grants read and write access to administrative information

//	@securitydefinitions.oauth2.implicit	OAuth2Implicit
//	@authorizationUrl						https://example.com/oauth/authorize
//	@scope.write							Grants write access
//	@scope.admin							Grants read and write access to administrative information

//	@securitydefinitions.oauth2.password	OAuth2Password
//	@tokenUrl								https://example.com/oauth/token
//	@scope.read								Grants read access
//	@scope.write							Grants write access
//	@scope.admin							Grants read and write access to administrative information

//	@securitydefinitions.oauth2.accessCode	OAuth2AccessCode
//	@tokenUrl								https://example.com/oauth/token
//	@authorizationUrl						https://example.com/oauth/authorize
//	@scope.admin							Grants read and write access to administrative information

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

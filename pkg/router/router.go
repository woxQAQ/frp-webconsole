package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(engine *gin.Engine) *Router {
	return &Router{engine: engine}
}

type Controller interface {
	Register(engine *gin.Engine)
}

func (r *Router) Register(controllers ...Controller) {
	for _, controller := range controllers {
		controller.Register(r.engine)
	}
}

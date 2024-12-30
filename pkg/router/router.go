package router

import (
	"github.com/gin-gonic/gin"
	"github.com/woxQAQ/frp-webconsole/pkg/controller"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(engine *gin.Engine) *Router {
	return &Router{engine: engine}
}
func (r *Router) Register(controllers ...controller.Controller) {
	for _, controller := range controllers {
		controller.Register(r.engine)
	}
}

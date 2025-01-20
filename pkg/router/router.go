package router

import (
	"github.com/go-fuego/fuego"
)

type Router struct {
	svc *fuego.Server
}

func NewRouter(svc *fuego.Server) *Router {
	return &Router{svc: svc}
}

type Controller interface {
	Register(svc *fuego.Server)
}

func (r *Router) Register(controllers ...Controller) {
	for _, controller := range controllers {
		controller.Register(r.svc)
	}
}

package controller

import (
	"github.com/gin-gonic/gin"
)

type FrpController struct{}

func NewFrpController() *FrpController {
	return &FrpController{}
}

func (f *FrpController) Register(engine *gin.RouterGroup) {
	engine.GET("/frp/config", f.GetFrpcConfig)
	engine.POST("/frp/install", f.InstallFrpc)
}

// GetFrpcConfig godoc
//
//	@Summary		Get Frpc Config
//	@Description	Get Frpc Config
//	@id				GetFrpcConfig
//	@Tags			frp
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/frp/config [get]
func (f *FrpController) GetFrpcConfig(ctx *gin.Context) {
}

// InstallFrpc godoc
//
//	@Summary		Install Frpc
//	@Description	Install Frpc
//	@id				InstallFrpc
//	@Tags			frp
//	@Accept			json
//	@Produce		json
//
//	@Success		200
//	@Router			/frp/install [post]
func (f *FrpController) InstallFrpc(ctx *gin.Context) {}

// ListFrpRelease godoc
//	@Summary		List Frp Release
//	@Description	List Frp Release
//	@id				ListFrpRelease
//	@Tags			frp
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Param			page		query	int				false	"page"
//	@Param			pageSize	query	int				false	"pageSize"
//	@Param			request		body	models.System	true	"SystemInfo"
//	@Router			/frp/release [get]
func (f *FrpController) ListFrpRelease(ctx *gin.Context) {}

package controller

import (
	"github.com/gin-gonic/gin"
)

type FrpControllerTemplate struct{}

func NewFrpController() *FrpControllerTemplate {
	return &FrpControllerTemplate{}
}

func (f *FrpControllerTemplate) Register(engine *gin.RouterGroup) {
	engine.GET("/frp/config", f.GetFrpcConfig)
	engine.POST("/frp/install", f.InstallFrpc)
}

// GetFrpcConfig godoc
//
// @Summary		Get Frpc Config
// @Description	Get Frpc Config
// @Tags			frp
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/frp/config [get]
func (f *FrpControllerTemplate) GetFrpcConfig(ctx *gin.Context) {}

// @Summary		Install Frpc
// @Description	Install Frpc
// @Tags			frp
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/frp/install [post]
func (f *FrpControllerTemplate) InstallFrpc(ctx *gin.Context) {}

// @Summary		List Frp Release
// @Description	List Frp Release
// @Tags		frp
// @Accept		json
// @Produce		json
// @Success		200
// @Param		page		query	int false	"page"
// @Param		pageSize	query	int false	"pageSize"
// @Param		request	body		models.System true "SystemInfo"
// @Router		/frp/release [get]
func (f *FrpControllerTemplate) ListFrpRelease(ctx *gin.Context) {}

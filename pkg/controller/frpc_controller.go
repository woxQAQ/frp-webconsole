package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FrpController struct {
}

func NewFrpController() *FrpController {
	return &FrpController{}
}

func (f *FrpController) Register(engine *gin.Engine) {
	engine.GET("/frp/config", f.GetFrpcConfig)
	engine.POST("/frp/install", f.InstallFrpc)
	engine.GET("/frp/release", f.ListFrpRelease)
}

// GetFrpcConfig godoc
//
//	@Summary		Get Frpc Config
//	@Description	Get Frpc Config
//	@Tags			frp
//	@Accept			json
//	@Produce		json
//	@Success		200
//	@Router			/frp/config [get]
func (f *FrpController) GetFrpcConfig(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

// @Summary		Install Frpc
// @Description	Install Frpc
// @Tags			frp
// @Accept			json
// @Produce		json
// @Success		200
// @Router			/frp/install [post]
func (f *FrpController) InstallFrpc(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
}

// @Summary		List Frp Release
// @Description	List Frp Release
// @Tags			frp
// @Accept			json
// @Produce		json
// @Success		200
// @Param			page		query	int				false	"page"
// @Param			pageSize	query	int				false	"pageSize"
// @Param			request		body	models.System	true	"SystemInfo"
// @Router			/frp/release [get]
func (f *FrpController) ListFrpRelease(ctx *gin.Context) {}

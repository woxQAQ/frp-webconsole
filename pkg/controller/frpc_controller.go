package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/woxQAQ/frp-webconsole/pkg/binding"
	httperrors "github.com/woxQAQ/frp-webconsole/pkg/errors/http"
	"github.com/woxQAQ/frp-webconsole/pkg/models"
	"github.com/woxQAQ/frp-webconsole/pkg/services"
)

type FrpController struct {
	service services.FrpcService
}

func NewFrpController(service services.FrpcService) *FrpController {
	return &FrpController{service: service}
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
//
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
func (f *FrpController) ListFrpRelease(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("pageSize", "10")
	system, err := binding.BindBody[models.System](ctx)
	if err != nil {
		httperrors.NewBadRequestError(ctx, err.Error())
		return
	}
	releases, err := f.service.ListFrpRelease(ctx.Request.Context(), cast.ToInt(page), cast.ToInt(pageSize), system)
	if err != nil {
		httperrors.NewBadRequestError(ctx, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, releases)
}

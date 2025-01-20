package controller

import (
	"net/http"

	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/option"

	"github.com/woxQAQ/frp-webconsole/pkg/models"
)

type FrpController struct {
}

func NewFrpController() *FrpController {
	return &FrpController{}
}

func (f *FrpController) Register(sv *fuego.Server) {
	fuego.Get(sv, "/frp/config", f.GetFrpcConfig)
	fuego.Post(sv, "/frp/install", f.InstallFrpc)
	fuego.Get(sv, "/frp/release", f.ListFrpRelease,
		option.Summary("List Frp Release"),
		option.Description("List Frp Release"),
		option.Tags("frp"),
		option.RequestContentType("application/json"),
		option.AddResponse(http.StatusNotFound, "not found", fuego.Response{
			ContentTypes: []string{"application/json"},
			Type:         fuego.HTTPError{},
		}),
	)
}

func (f *FrpController) GetFrpcConfig(ctx fuego.ContextWithBody[string]) (string, error) {
	return "Hello, World!", nil
}

func (f *FrpController) InstallFrpc(ctx fuego.ContextWithBody[string]) (string, error) {
	return "Hello, World!", nil
}

func (f *FrpController) ListFrpRelease(ctx fuego.ContextWithBody[models.System]) ([]models.FrpRelease, error) {
	return []models.FrpRelease{}, nil
}

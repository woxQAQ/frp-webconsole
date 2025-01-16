package controllers

import (
	"context"

	"github.com/google/go-github/v68/github"
	"github.com/woxQAQ/frp-webconsole/pkg/gen/frpc"
	"github.com/woxQAQ/frp-webconsole/pkg/stores"

	frpcSvc "github.com/woxQAQ/frp-webconsole/pkg/gen/http/frpc/server"
	"github.com/woxQAQ/frp-webconsole/pkg/services"
	goahttp "goa.design/goa/v3/http"
)

func NewMux(ctx context.Context) goahttp.Muxer {
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)
	mux := goahttp.NewMuxer()
	// mux.Use(middleware.Logger())
	// stores declaration
	ghClient := github.NewClient(nil)
	ghStore := stores.NewGithubClient(ghClient)

	// services declaration
	frpcService := services.NewFrpcService(ghStore)

	// endpoints declaration
	endpoints := frpc.NewEndpoints(frpcService)

	// controllers declaration
	frpcsvc := frpcSvc.New(endpoints, mux, dec, enc, nil, nil)
	frpcSvc.Mount(mux, frpcsvc)
	return mux
}

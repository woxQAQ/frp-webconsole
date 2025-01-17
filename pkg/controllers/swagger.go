package controllers

import (
	"encoding/json"
	"net/http"
	"path"
	"runtime"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/handlers"
)

func SwaggerHandler() http.Handler {
	_, file, _, _ := runtime.Caller(0)
	specDoc, err := loads.Spec(path.Join(path.Dir(file), "../gen/http/openapi.json"))
	if err != nil {
		panic(err)
	}
	base := "/"
	handler := http.NotFoundHandler()
	handler = middleware.SwaggerUI(middleware.SwaggerUIOpts{
		BasePath: base,
		Path:     "docs",
		SpecURL:  "/docs/swagger.json",
	}, handler)

	b, err := json.MarshalIndent(specDoc.Spec(), "", "  ")
	if err != nil {
		panic(err)
	}

	handler = handlers.CORS()(middleware.Spec("/docs", b, handler))

	return handler
}

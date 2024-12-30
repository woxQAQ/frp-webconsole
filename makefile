API_DIR = api

.PHONY: swag
swag:
	@type -p swag || go install github.com/swaggo/swag/cmd/swag@latest
	@type -p openapi-generator || go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	@type -p swagger2openapi || npm install -g swagger2openapi

.PHONY: api
api: swag
	swag init -g ${API_DIR}/meta.go -o ${API_DIR}
	swagger2openapi -y ${API_DIR}/swagger.yaml > ${API_DIR}/openapi_bundle.yaml
	oapi-codegen -package api -generate types -o ${API_DIR}/types.gen.go ${API_DIR}/openapi_bundle.yaml
	oapi-codegen -package api -generate gin -o ${API_DIR}/server.gen.go ${API_DIR}/openapi_bundle.yaml
	rm -rf ${API_DIR}/openapi_bundle.yaml

.PHONY: fmt
fmt:
	go fmt ./...
	go vet ./...

.PHONY: run
run: fmt
	go run cmd/main.go

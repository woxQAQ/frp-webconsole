API_DIR = api

.PHONY: swag
swag:
	@type -p swag || go install github.com/swaggo/swag/cmd/swag@latest
	@type -p openapi-generator || go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	@type -p swagger2openapi || npm install -g swagger2openapi

.PHONY: api
api: swag
	swag init -g ${API_DIR}/meta.go -o ${API_DIR}
	swag fmt
.PHONY: fmt
fmt:
	go fmt ./...
	go vet ./...

.PHONY: run
run: fmt
	go run cmd/main.go

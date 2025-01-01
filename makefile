API_DIR = api

.PHONY: swag
swag:
	@type -p swag || go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: api
api: swag
	swag init -g ${API_DIR}/meta.go -o ${API_DIR}
	swag fmt
.PHONY: fmt
fmt:
	go fmt ./...
	go vet ./...
	swag fmt

.PHONY: module
module:
	go mod tidy --compat=1.23
	go mod verify

.PHONY: run
run: fmt
	go run cmd/main.go

API_DIR = api

.PHONY: api
api:
	swag init -g ${API_DIR}/meta.go -o ${API_DIR}

.PHONY: fmt
fmt:
	go fmt ./...
	go vet ./...

.PHONY: run
run: fmt
	go run cmd/main.go

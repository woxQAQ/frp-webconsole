API_DIR = api

.PHONY: fmt
fmt:
	go fmt ./...
	go vet ./...

.PHONY: module
module:
	go mod tidy --compat=1.23
	go mod verify

.PHONY: run
run: fmt
	go run cmd/main.go

.PHONY: module
module: fmt
	go mod tidy
	go vet ./...
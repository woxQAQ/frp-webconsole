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

.PHONY: module
module: fmt
	go mod tidy
	go mod verify

GO_FILES := $(shell git ls-files | grep "\.go$$" | grep -v api)

.PHONY: imports
imports:
	for file in $(GO_FILES); do \
		bash hack/merge-imports.sh $$file; \
	done; \
	goimports -local github.com/woxQAQ/frp-webconsole -w $(GO_FILES)

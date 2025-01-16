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

.PHONY: gen
gen: 
	goa gen github.com/woxQAQ/frp-webconsole/api 
	goa example github.com/woxQAQ/frp-webconsole/api -o example PACKAGE=github.com/woxQAQ/frp-webconsole/pkg/gen

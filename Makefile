#Docker shortcuts
#Go shortcuts
.PHONY:all
all: fmt lint build

.PHONY:build
build:
	go build .

.PHONY:fmt
fmt:
	gofmt -s -w .

.PHONY:lint
lint:
	golangci-lint run ./...

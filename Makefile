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

.PHONY:test
test:
	go run . --csv=test_data/test_01.csv --filter=top_by_category

export CGO_ENABLED=0

all: build

.PHONY: build
build: modules
	GOOS=linux GOARCH=amd64 go build -o bin/ ./cmd/conqueror

.PHONY: modules
modules:
	go mod tidy

.PHONY: test
test:
	go test ./...

.PHONY: check
check:
	golangci-lint run

.PHONY: clean
clean:
	rm -rf bin/conqueror

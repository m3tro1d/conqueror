all: build test check

.PHONY: modules
modules:
	go mod tidy

.PHONY: build
build: modules
	go build -o bin/ ./cmd/conqueror

.PHONY: clean
clean:
	rm -rf bin

.PHONY: test
test:
	go test ./...

.PHONY: check
check:
	golangci-lint run

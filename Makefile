GOFILES := $(shell find . -name '*.go' -not -path './vendor/*')

.PHONY: test fmt vet lint upgrade

test:
	go test ./... -timeout 15s -cover -coverprofile=coverage.out -v
	go tool cover -html=coverage.out -o coverage.html

fmt:
	gofmt -w $(GOFILES)

vet:
	go vet ./...

lint:
	@command -v golangci-lint >/dev/null 2>&1 || { echo "golangci-lint not installed. Install it from https://golangci-lint.run/usage/install/ to run lint."; exit 1; }
	golangci-lint run ./...

upgrade:
	go get -u all
	go mod tidy

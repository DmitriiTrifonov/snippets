.PHONY: build lint

bin:
	mkdir bin

install-lint: bin
	GOBIN=${PWD}/bin/ go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.48.0

build:
	go build -o bin/snippets cmd/snippets/main.go

lint: install-lint
	bin/golangci-lint run ./...
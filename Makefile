.DEFAULT_GOAL := run

fmt:
	go fmt ./...
.PHONY:fmt

vet: fmt
	go vet ./...
.PHONY:vet

run: vet
	go run ./cmd/cli/main.go
.PHONY: run

clean:
	go clean
.PHONY:clean
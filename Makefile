.DEFAULT_GOAL := run

fmt:
	go fmt ./...
.PHONY:fmt

vet: fmt
	go vet ./...
.PHONY:vet

run: vet
	go run ./cmd/learning-go/main.go
.PHONY: run

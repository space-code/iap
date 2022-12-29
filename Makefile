all: imports fmt lint vet

fmt:
	go fmt ./...
.PHONY:fmt

lint:
	golint ./...
.PHONY:lint

vet:
	go vet -v ./...
.PHONY:vet

imports:
	goimports -l -w .
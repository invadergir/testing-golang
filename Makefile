.DEFAULT_GOAL := build

#.PHONY:fmt vet build
clean:
	rm -fv bin/*

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: clean vet
	go build -o bin/testing-go ./...

test: build
	go test ./...

run: build
	bin/testing-go


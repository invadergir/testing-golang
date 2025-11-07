.DEFAULT_GOAL := build

#.PHONY:fmt vet build
clean:
	rm -fv bin/*

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: clean vet
	#go build -o bin/testing-golang ./...
	#go build -o bin/testing-golang .
	go build  ./...

test: build
	go test ./...

testv: build
	go test ./... -v

run: build
	bin/testing-golang


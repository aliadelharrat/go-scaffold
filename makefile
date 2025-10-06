
.DEFAULT_GOAL := build

.PHONY:clean fmt vet build

clean:
	go clean
fmt: clean
	go fmt ./...
vet: fmt
	go vet ./...
build: vet
	go build -o main

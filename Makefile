BINARY := storage-app


BINARY := item-app

help:
	@echo "usage: make <command>"
	@echo
	@echo "commands:"
	@echo "	build	compiles and build the app"
	@echo "	start	starts the app"

start:
	@./bin/${BINARY}

build:
	@go build -o bin/${BINARY} cmd/main.go
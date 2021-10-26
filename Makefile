BINARY := storage-app

build:
	@go build -o bin/${BINARY} cmd/main.go
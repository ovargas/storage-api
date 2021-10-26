# storage-api

A simple API created as a PoC to play a bit with gRCP and protocol buffer

Build with Makefile

```bash
make
```

Build using `go`

```bash
go build -o bin/storage-app cmd/main.go
```

Execute the service

```bash
./bin/storage-app
```

Display command line options

```bash
./bin/storage-app --help
```

```text
Usage of ./storage-app:
  -port int
        The server port (default 10000)
  -storage_folder string
        The folder to store the files (default "tmp")
```

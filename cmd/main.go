package main

import (
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	pb "github.com/ovargas/api-go/storage/v1"
	_ "github.com/ovargas/storage-api/file_system"
	"github.com/ovargas/storage-api/internal/server"
	"github.com/ovargas/storage-api/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var (
	port = flag.Int("port", 10000, "The server port")
	folder = flag.String("storage_folder", "tmp", "The folder to store the files")
)

func main() {
	flag.Parse()
	// Creating a storage provider to be used with the service
	store, err := storage.New("FileSystem", map[string]interface{}{
		"folder": *folder,
	})

	//Configuring the logging
	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logEntry := logrus.NewEntry(logger)
	grpc_logrus.ReplaceGrpcLogger(logEntry)

	//Configuring logrus for gRCP
	opts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	//Creating a gRCP server and registering the middlewares
	grpcServer := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_logrus.UnaryServerInterceptor(logEntry, opts...),
	))

	// Registering the storage service
	pb.RegisterStorageServiceServer(grpcServer, server.New(store))

	if err != nil {
		log.Fatalf("unable to create storage: %v", err)
	}

	// The TCP listener where the service will be allocated
	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	logger.Info("starting the service in the port ", *port)

	//Start the server
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("unable to start server: %v", err)
	}
}

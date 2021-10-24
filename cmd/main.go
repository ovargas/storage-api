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
)

func main() {

	listen, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	store, err := storage.New("FileSystem", map[string]interface{}{
		"folder": "tmp",
	})

	logger := logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logEntry := logrus.NewEntry(logger)
	grpc_logrus.ReplaceGrpcLogger(logEntry)

	opts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	grpcServer := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		grpc_logrus.UnaryServerInterceptor(logEntry, opts...),
	))

	pb.RegisterStorageServiceServer(grpcServer, server.New(store))

	if err != nil {
		log.Fatalf("unable to create storage: %v", err)
	}

	logger.Info("server start")
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatalf("unable to start server: %v", err)
	}
}

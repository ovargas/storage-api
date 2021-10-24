package server

import (
	"context"
	pb "github.com/ovargas/api-go/storage/v1"
	"github.com/ovargas/storage-api/storage"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	pb.UnimplementedStorageServiceServer
	storage *storage.Storage
}

func New(storage *storage.Storage) *server {
	return &server{
		storage: storage,
	}
}

func (s *server) Download(ctx context.Context, request *pb.DownloadRequest) (*pb.File, error) {
	bytes, err := s.storage.Download(ctx, request.Filename)
	if err != nil {
		return nil, err
	}

	return &pb.File{
		Name:    request.Filename,
		Content: &pb.File_Bytes{Bytes: bytes},
	}, nil
}

func (s *server) Create(ctx context.Context, request *pb.CreateRequest) (*pb.File, error) {
	url, err := s.storage.Store(ctx, request.Filename, request.GetBytes())
	if err != nil {
		return nil, err
	}
	return &pb.File{
		Name:      request.Filename,
		MediaType: request.MediaType,
		Content:   &pb.File_Url{Url: url},
	}, nil
}

func (s *server) Delete(ctx context.Context, request *pb.DeleteRequest) (*emptypb.Empty, error) {
	err := s.storage.Delete(ctx, request.Filename)
	return &emptypb.Empty{}, err
}

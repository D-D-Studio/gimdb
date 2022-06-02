package grpc

import (
	"context"
	"gimdb/domain"
	"gimdb/infrastructure/rpc/grpc"
	"log"
)

type StorageService struct {
	grpc.UnimplementedStorageServiceServer

	workspace *domain.Workspace
}

func NewStorageService(workspace *domain.Workspace) *StorageService {
	storageService := new(StorageService)

	storageService.workspace = workspace

	return storageService
}

func (s *StorageService) CreateStorage(ctx context.Context, request *grpc.CreateStorageRequest) (*grpc.CreateStorageResponse, error) {
	log.Printf("[gRPC] CREATE_STORAGE \"%s\"\n", request.Name)

	return &grpc.CreateStorageResponse{}, s.workspace.CreateStorage(request.Name)
}

func (s *StorageService) DeleteStorage(ctx context.Context, request *grpc.DeleteStorageRequest) (*grpc.DeleteStorageResponse, error) {
	log.Printf("[gRPC] DELETE_STORAGE \"%s\"\n", request.Name)

	return &grpc.DeleteStorageResponse{}, s.workspace.DeleteStorage(request.Name)
}

func (s *StorageService) ExistsStorage(ctx context.Context, request *grpc.ExistsStorageRequest) (*grpc.ExistsStorageResponse, error) {
	log.Printf("[gRPC] EXISTS_STORAGE \"%s\"\n", request.Name)

	return &grpc.ExistsStorageResponse{IsExists: s.workspace.IsStorageExists(request.Name)}, nil
}

func (s *StorageService) Get(ctx context.Context, request *grpc.GetRequest) (*grpc.GetResponse, error) {
	log.Printf("[gRPC] GET \"%s\" \"%s\"\n", request.Storage, request.Key)

	storage, err := s.workspace.GetStorage(request.Storage)

	if err != nil {
		return nil, err
	}

	value, err := storage.Get(request.Key)

	if err != nil {
		return nil, err
	}

	return &grpc.GetResponse{Value: value}, nil
}

func (s *StorageService) Set(ctx context.Context, request *grpc.SetRequest) (*grpc.SetResponse, error) {
	log.Printf("[gRPC] SET \"%s\" \"%s\" \"%s\"\n", request.Storage, request.Key, request.Value)

	storage, err := s.workspace.GetStorage(request.Storage)

	if err != nil {
		return nil, err
	}

	storage.Set(request.Key, request.Value)

	return &grpc.SetResponse{}, nil
}

func (s *StorageService) Delete(ctx context.Context, request *grpc.DeleteRequest) (*grpc.DeleteResponse, error) {
	log.Printf("[gRPC] DELETE \"%s\" \"%s\"\n", request.Storage, request.Key)

	storage, err := s.workspace.GetStorage(request.Storage)

	if err != nil {
		return nil, err
	}

	return &grpc.DeleteResponse{}, storage.Delete(request.Key)
}

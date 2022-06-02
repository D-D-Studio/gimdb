package rpc

import (
	api_grpc "gimdb/api/grpc"
	infra_grpc "gimdb/infrastructure/rpc/grpc"
	proto_grpc "google.golang.org/grpc"
	"log"
	"net"
)

type GrpcServer struct {
	handler *proto_grpc.Server
}

func NewGrpcServer(storageService *api_grpc.StorageService) *GrpcServer {
	server := new(GrpcServer)

	server.handler = proto_grpc.NewServer()
	infra_grpc.RegisterStorageServiceServer(server.handler, storageService)

	return server
}

func (s *GrpcServer) Listen(addr string) {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	if err := s.handler.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}

package main

import (
	"flag"
	g_rpc "gimdb/api/grpc"
	json_rpc "gimdb/api/json-rpc"
	"gimdb/domain"
	"gimdb/infrastructure/rpc"
	"log"
	"strconv"
)

var (
	jsonRpc     = flag.Bool("json-rpc", false, "Enable JSON-RPC server")
	jsonRpcHost = flag.String("json-rpc-host", "0.0.0.0", "JSON-RPC host")
	jsonRpcPort = flag.Int("json-rpc-port", 6543, "JSON-RPC port")
	gRpcHost    = flag.String("grpc-host", "0.0.0.0", "gRPC host")
	gRpcPort    = flag.Int("grpc-port", 6544, "gRPC port")
)

func main() {
	flag.Parse()
	channel := make(chan int)

	log.Println("Initialize database...")

	workspace := domain.NewWorkspace()

	log.Println("Initialize gRPC server...")

	gRpcStorageService := g_rpc.NewStorageService(workspace)
	gRpcServer := rpc.NewGrpcServer(gRpcStorageService)

	log.Println("Starting gRPC server...")

	go gRpcServer.Listen(*gRpcHost + ":" + strconv.Itoa(*gRpcPort))

	log.Printf("gRPC server started at %s\n", *gRpcHost+":"+strconv.Itoa(*gRpcPort))

	if *jsonRpc {
		log.Println("Initialize JSON-RPC server...")

		jsonRpcStorageService := json_rpc.NewStorageService(workspace)
		jsonRpcServer := rpc.NewJsonRpcServer(jsonRpcStorageService)

		log.Println("Starting JSON-RPC server...")

		go jsonRpcServer.Listen(*jsonRpcHost + ":" + strconv.Itoa(*jsonRpcPort))

		log.Printf("JSON-RPC server started at %s\n", *jsonRpcHost+":"+strconv.Itoa(*jsonRpcPort))
	}

	<-channel
}

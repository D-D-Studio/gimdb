package rpc

import (
	json_rpc "gimdb/api/json-rpc"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"log"
	"net/http"
)

type JsonRpcServer struct {
	handler *rpc.Server
}

func NewJsonRpcServer(storageService *json_rpc.StorageService) *JsonRpcServer {
	server := new(JsonRpcServer)

	server.handler = rpc.NewServer()
	server.handler.RegisterCodec(json2.NewCodec(), "application/json")
	err := server.handler.RegisterService(storageService, "Storage")

	if err != nil {
		log.Fatalf("cannot register service %v\n", err)
	}

	return server
}

func (j *JsonRpcServer) Listen(addr string) {
	log.Fatalln(http.ListenAndServe(addr, j.handler))
}

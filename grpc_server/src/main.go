package main

import (
	"github.com/maei/golang_grpc_unary_errorhandling/grpc_server/src/server"
	"github.com/maei/shared_utils_go/logger"
)

func main() {
	logger.Info("gRPC-Server: Start gRPC-Server started")
	server.StartGRPCServer()
}

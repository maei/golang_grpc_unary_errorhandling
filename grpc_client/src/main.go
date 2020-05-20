package main

import (
	"github.com/maei/golang_grpc_unary_errorhandling/grpc_client/src/service"
	"github.com/maei/shared_utils_go/logger"
)

func main() {
	logger.Info("gRPC-Client: Start gRPC-Client started")
	service.SquareService.GetSquare()
}

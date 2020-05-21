package main

import (
	"github.com/maei/golang_grpc_unary_errorhandling/grpc_client/src/client"
	"github.com/maei/golang_grpc_unary_errorhandling/grpc_client/src/domain/squarepb"
	"github.com/maei/golang_grpc_unary_errorhandling/grpc_client/src/service"
	"github.com/maei/shared_utils_go/logger"
	"time"
)

func main() {
	logger.Info("gRPC-Client: Start gRPC-Client started")
	conn, connErr := client.GRPCClient.SetClient()
	if connErr != nil {
		logger.Error("gRPC-Client: Error while creating connection obj", connErr)
	}
	defer conn.Close()
	c := squarepb.NewSquareRootServiceClient(conn)

	service.SquareService.GetSquare(c, 5*time.Second)
}

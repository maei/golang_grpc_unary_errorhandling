package service

import (
	"context"
	"fmt"
	"github.com/maei/golang_grpc_unary_errorhandling/grpc_client/src/client"
	"github.com/maei/golang_grpc_unary_errorhandling/grpc_client/src/domain/squarepb"
	"github.com/maei/shared_utils_go/logger"
)

var SquareService squareServiceInterface = &squareService{}

type squareServiceInterface interface {
	GetSquare()
}

type squareService struct{}

func (*squareService) GetSquare() {
	conn, connErr := client.GRPCClient.SetClient()
	if connErr != nil {
		logger.Error("gRPC-Client: Error while creating connection obj", connErr)
	}
	c := squarepb.NewSquareRootServiceClient(conn)

	req := &squarepb.SquareRootRequest{
		A: float32(4),
	}

	res, resError := c.GetSquareRoot(context.Background(), req)
	if resError != nil {
		logger.Error("gRPC-Client: Error while receiving request from Server", resError)
	}
	fmt.Printf("gRPC-Client: Result from Square-Root computition %v", res.GetResult())
}

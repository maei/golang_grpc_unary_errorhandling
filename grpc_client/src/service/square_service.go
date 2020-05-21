package service

import (
	"context"
	"fmt"
	"github.com/maei/golang_grpc_unary_errorhandling/grpc_client/src/client"
	"github.com/maei/golang_grpc_unary_errorhandling/grpc_client/src/domain/squarepb"
	"github.com/maei/shared_utils_go/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var SquareService squareServiceInterface = &squareService{}

type squareServiceInterface interface {
	GetSquare(duration time.Duration)
}

type squareService struct{}

func (*squareService) GetSquare(duration time.Duration) {
	conn, connErr := client.GRPCClient.SetClient()
	if connErr != nil {
		logger.Error("gRPC-Client: Error while creating connection obj", connErr)
	}
	c := squarepb.NewSquareRootServiceClient(conn)

	req := &squarepb.SquareRootRequest{
		A: float32(16),
	}

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	res, resError := c.GetSquareRoot(ctx, req)
	if resError != nil {
		grpcErr, ok := status.FromError(resError)
		if ok {
			if grpcErr.Code() == codes.DeadlineExceeded {
				logger.Error("gRPC-Client Error: Deadline Exceeded", grpcErr.Err())
				return
			}
			if grpcErr.Code() == codes.InvalidArgument {
				logger.Error("gRPC-Client Error: Invalid Argument", grpcErr.Err())
				return
			} else {
				logger.Error("gRPC-Client unexpected Error", grpcErr.Err())
			}
		} else {
			logger.Error("gRPC-Client: Error while receiving request from Server", resError)
			return
		}

	}
	fmt.Printf("gRPC-Client: Result from Square-Root computition %v", res.GetResult())
}

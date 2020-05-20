package server

import (
	"context"
	"fmt"
	"github.com/maei/golang_grpc_unary_errorhandling/grpc_server/src/domain/squarepb"
	"github.com/maei/shared_utils_go/logger"
	"google.golang.org/grpc"
	"math"
	"net"
)

type server struct{}

var (
	s = grpc.NewServer()
)

func (*server) GetSquareRoot(ctx context.Context, req *squarepb.SquareRootRequest) (*squarepb.SquareRootResponse, error) {
	logger.Info(fmt.Sprintf("Getting gRPC-Client request with: %v", req.GetA()))
	a := req.GetA()
	squareA := math.Sqrt(float64(a))

	res := &squarepb.SquareRootResponse{
		Result: float32(squareA),
	}

	return res, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Error("error while listening gRPC Server", err)
	}

	squarepb.RegisterSquareRootServiceServer(s, &server{})

	errServer := s.Serve(lis)
	if errServer != nil {
		logger.Error("error while serve gRPC Server", errServer)
	}
}

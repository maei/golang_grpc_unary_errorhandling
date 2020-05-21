package server

import (
	"context"
	"fmt"
	"github.com/maei/golang_grpc_unary_errorhandling/grpc_server/src/domain/squarepb"
	"github.com/maei/shared_utils_go/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math"
	"net"
	"time"
)

type server struct{}

var (
	s = grpc.NewServer()
)

func (*server) GetSquareRoot(ctx context.Context, req *squarepb.SquareRootRequest) (*squarepb.SquareRootResponse, error) {
	logger.Info(fmt.Sprintf("Getting gRPC-Client request with: %v", req.GetA()))
	/*	for i := 0; i < 3; i++ {


		time.Sleep(1 * time.Second)
	}*/
	check := ctx.Err()

	switch check {
	case context.DeadlineExceeded:
		logger.Info("Client deadline exceeded!")
		return nil, status.Errorf(codes.DeadlineExceeded, "Client deadline exceeded!")
	case context.Canceled:
		logger.Info("Client canceled RPC!")
		return nil, status.Errorf(codes.Canceled, "The client canceled the request!")
	default:
		fmt.Println(check)
	}

	a := req.GetA()

	if math.Signbit(float64(a)) {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("gRPC-Server: Request should be a positiv value. Input %v", a))
	}
	squareA := math.Sqrt(float64(a))
	time.Sleep(3 * time.Second)

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

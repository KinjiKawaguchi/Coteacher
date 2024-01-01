package grpc

import (
	"context"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/healthcheck"

	pb "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"
)

type healthcheckServiceServer struct {
	healthcheckInteractor *healthcheck.Interactor
}

func NewHealthcheckServiceServer() pb.HealthcheckServiceServer {
	return &healthcheckServiceServer{
		healthcheckInteractor: healthcheck.NewInteractor(),
	}
}

func (i *healthcheckServiceServer) Unary(ctx context.Context, req *pb.UnaryRequest) (*pb.UnaryResponse, error) {
	return i.healthcheckInteractor.Unary(ctx, req)
}

func (i *healthcheckServiceServer) ServerStreaming(req *pb.ServerStreamingRequest, stream pb.HealthcheckService_ServerStreamingServer) error {
	return i.healthcheckInteractor.ServerStreaming(req, stream)
}

func (i *healthcheckServiceServer) ClientStreaming(stream pb.HealthcheckService_ClientStreamingServer) error {
	return i.healthcheckInteractor.ClientStreaming(stream)
}

func (i *healthcheckServiceServer) BidirectionalStreaming(stream pb.HealthcheckService_BidirectionalStreamingServer) error {
	return i.healthcheckInteractor.BidirectionalStreaming(stream)
}

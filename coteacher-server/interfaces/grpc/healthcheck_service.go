package grpc

import (
	"context"

	"connectrpc.com/connect"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/healthcheck"

	coteacherv1 "github.com/KinjiKawaguchi/Coteacher/coteacher-server/proto-gen/go/coteacher/v1"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/proto-gen/go/coteacher/v1/coteacherv1connect"
)

type healthcheckServiceServer struct {
	healthcheckInteractor *healthcheck.Interactor
}

func NewHealthCheckServiceServer(interactor *healthcheck.Interactor) coteacherv1connect.HealthcheckServiceHandler {
	return &healthcheckServiceServer{interactor}
}

func (s *healthcheckServiceServer) Ping(ctx context.Context, req *connect.Request[coteacherv1.PingRequest]) (*connect.Response[coteacherv1.PingResponse], error) {
	resp := s.healthcheckInteractor.Ping(req.Msg)
	return connect.NewResponse(resp), nil
}

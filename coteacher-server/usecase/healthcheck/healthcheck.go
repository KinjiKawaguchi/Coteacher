package healthcheck

import (
	pb "github.com/KinjiKawaguchi/Coteacher/coteacher-server/proto-gen/go/coteacher/v1"
)

type Interactor struct{}

func NewInteractor() *Interactor {
	return &Interactor{}
}

func (i *Interactor) Ping(req *pb.PingRequest) *pb.PingResponse {
	return &pb.PingResponse{}
}

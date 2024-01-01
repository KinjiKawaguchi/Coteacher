package healthcheck

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"
)

type Interactor struct{}

func NewInteractor() *Interactor {
	return &Interactor{}
}

func (*Interactor) Unary(ctx context.Context, req *pb.UnaryRequest) (*pb.UnaryResponse, error) {
	return &pb.UnaryResponse{
		Msg: "🐔 < " + req.Msg,
	}, nil
}

func (*Interactor) ServerStreaming(req *pb.ServerStreamingRequest, stream pb.HealthcheckService_ServerStreamingServer) error {
	i := 1
	for {
		if err := stream.Send(&pb.ServerStreamingResponse{
			Msg: fmt.Sprintf("🐔(%d) < %s", i, req.Msg),
		}); err != nil {
			log.Println(err)
			return err
		}
		time.Sleep(time.Second)
		i++
	}
}

func (*Interactor) ClientStreaming(stream pb.HealthcheckService_ClientStreamingServer) error {
	msgs := make([]string, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return err
		}
		msgs = append(msgs, req.Msg)
	}
	return stream.SendAndClose(&pb.ClientStreamingResponse{
		Msg: fmt.Sprintf("🐔 < %v", msgs),
	})
}

func (*Interactor) BidirectionalStreaming(stream pb.HealthcheckService_BidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			return err
		}
		if err := stream.Send(&pb.BidirectionalStreamingResponse{
			Msg: fmt.Sprintf("🐔 < %s", req.Msg),
		}); err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

var _ pb.HealthcheckServiceServer = (*Interactor)(nil)

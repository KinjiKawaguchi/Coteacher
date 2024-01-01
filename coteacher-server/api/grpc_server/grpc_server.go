package grpc_server

import (
	"context"
	"log"

	grpc_interfaces "github.com/KinjiKawaguchi/Coteacher/coteacher-server/interfaces/grpc"

	pb "github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func New(opts ...optionFunc) *grpc.Server {
	opt := defaultOption()
	for _, f := range opts {
		f(opt)
	}

	serverOptions := make([]grpc.ServerOption, 0)
	// set logging interceptor
	serverOptions = append(serverOptions,
		grpc.ChainUnaryInterceptor(logging.UnaryServerInterceptor(interceptorLogger(opt.logger))),
		grpc.ChainStreamInterceptor(logging.StreamServerInterceptor(interceptorLogger(opt.logger))),
	)

	log.Println("Starting gRPC server...[grpc_server.go]")

	srv := grpc.NewServer(serverOptions...)
	reflection.Register(srv) // for grpcui
	healthcheckSrv := grpc_interfaces.NewHealthcheckServiceServer()
	pb.RegisterHealthcheckServiceServer(srv, healthcheckSrv)
	userSrv := grpc_interfaces.NewUserServiceServer(opt.entClient, opt.logger)
	pb.RegisterUserServiceServer(srv, userSrv)
	classSrv := grpc_interfaces.NewClassServiceServer(opt.entClient, opt.logger)
	pb.RegisterClassServiceServer(srv, classSrv)
	classInvitationCodeSrv := grpc_interfaces.NewClassInvitationCodeServiceServer(opt.entClient, opt.logger)
	pb.RegisterClassInvitationCodeServiceServer(srv, classInvitationCodeSrv)
	studentClassSrv := grpc_interfaces.NewStudentClassServiceServer(opt.entClient, opt.logger)
	pb.RegisterStudentClassServiceServer(srv, studentClassSrv)

	log.Println("gRPC server started. [grpc_server.go]")

	return srv
}

func interceptorLogger(l *slog.Logger) logging.Logger {
	log.Println("interceptorLogger [grpc_server.go]")
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}

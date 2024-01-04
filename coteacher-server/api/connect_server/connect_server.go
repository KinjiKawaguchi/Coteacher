package connect_server

import (
	"net/http"

	"connectrpc.com/connect"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/api/connect_server/interceptor"
	grpc_interfaces "github.com/KinjiKawaguchi/Coteacher/coteacher-server/interfaces/grpc"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/class"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/class_invitation_code"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/student_class"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/usecase/user"
	"github.com/rs/cors"

	"github.com/KinjiKawaguchi/Coteacher/proto-gen/go/coteacher/v1/coteacherv1connect"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func New(addr string, opts ...optionFunc) *http.Server {
	opt := defaultOption()
	for _, f := range opts {
		f(opt)
	}

	interceptors := connect.WithInterceptors(
		interceptor.Logger(),
	)

	mux := http.NewServeMux()

	classinvitationcodeSrv := grpc_interfaces.NewClassInvitationCodeServiceServer(class_invitation_code.NewInteractor(opt.entClient, opt.logger))
	mux.Handle(coteacherv1connect.NewClassInvitationCodeServiceHandler(classinvitationcodeSrv, interceptors))
	classSrv := grpc_interfaces.NewClassServiceServer(class.NewInteractor(opt.entClient, opt.logger))
	mux.Handle(coteacherv1connect.NewClassServiceHandler(classSrv, interceptors))
	studentclassSrv := grpc_interfaces.NewStudentClassServiceServer(student_class.NewInteractor(opt.entClient, opt.logger))
	mux.Handle(coteacherv1connect.NewStudentClassServiceHandler(studentclassSrv, interceptors))
	userSrv := grpc_interfaces.NewUserServiceServer(user.NewInteractor(opt.entClient, opt.logger))
	mux.Handle(coteacherv1connect.NewUserServiceHandler(userSrv, interceptors))

	corsHandler := cors.New(cors.Options{
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedOrigins: []string{"http://localhost:5000", "http://localhost:3000", opt.FrontendURL},
		AllowedHeaders: []string{
			"Accept-Encoding",
			"Content-Encoding",
			"Content-Type",
			"Connect-Protocol-Version",
			"Connect-Timeout-Ms",
			"Connect-Accept-Encoding",  // Unused in web browsers, but added for future-proofing
			"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
			"Grpc-Timeout",             // Used for gRPC-web
			"X-Grpc-Web",               // Used for gRPC-web
			"X-User-Agent",             // Used for gRPC-web
			"Authorization",
		},
		ExposedHeaders: []string{
			"Content-Encoding",         // Unused in web browsers, but added for future-proofing
			"Connect-Content-Encoding", // Unused in web browsers, but added for future-proofing
			"Grpc-Status",              // Required for gRPC-web
			"Grpc-Message",             // Required for gRPC-web
		},
	})

	return &http.Server{
		Addr: addr,
		Handler: h2c.NewHandler(
			corsHandler.Handler(mux),
			&http2.Server{},
		),
	}
}

module github.com/KinjiKawaguchi/Coteacher/coteacher-server

replace github.com/KinjiKawaguchi/Coteacher/proto-gen/go => ../proto-gen/go

go 1.21.5

require (
	connectrpc.com/grpcreflect v1.2.0
	entgo.io/ent v0.12.5
	github.com/google/uuid v1.3.1
	github.com/rs/cors v1.7.0
)

require (
	ariga.io/atlas v0.16.0 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	golang.org/x/mod v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240102182953-50ed04b92917 // indirect
)

require (
	connectrpc.com/connect v1.14.0
	github.com/KinjiKawaguchi/Coteacher/proto-gen/go v0.0.0-20240101035838-670236f5297b
	github.com/go-sql-driver/mysql v1.7.1
	github.com/joho/godotenv v1.5.1
	github.com/kelseyhightower/envconfig v1.4.0
	golang.org/x/exp v0.0.0-20231226003508-02704c960a9b
	golang.org/x/net v0.19.0
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/grpc v1.60.1 //indirect
	google.golang.org/protobuf v1.32.0
)

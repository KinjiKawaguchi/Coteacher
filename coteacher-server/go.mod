module github.com/KinjiKawaguchi/Coteacher/coteacher-server

replace github.com/KinjiKawaguchi/Coteacher/coteacher-server/proto-gen/go => ./proto-gen/go

go 1.21.5

require (
	connectrpc.com/connect v1.15.0
	connectrpc.com/grpcreflect v1.2.0
	entgo.io/ent v0.13.1
	github.com/KinjiKawaguchi/Coteacher/coteacher-server/proto-gen/go v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.7.1
	github.com/google/uuid v1.6.0
	github.com/joho/godotenv v1.5.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/rs/cors v1.10.1
	golang.org/x/exp v0.0.0-20240222234643-814bf88cf225
	golang.org/x/net v0.21.0
	google.golang.org/protobuf v1.32.0
)

require (
	ariga.io/atlas v0.19.1-0.20240203083654-5948b60a8e43 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	golang.org/x/mod v0.15.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240123012728-ef4313101c80 // indirect
	google.golang.org/grpc v1.62.0 // indirect
)

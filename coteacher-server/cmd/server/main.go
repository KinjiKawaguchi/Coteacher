package main

import (
	"coteacher/usecase/user" // gRPCサービスのパッケージ
	"database/sql"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var db *sql.DB

func main() {
	// .envファイルの読み込み、データベース接続など
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	db, err = sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}

	// gRPCサーバの設定
	lis, err := net.Listen("tcp", ":50051") // 50051はgRPCのデフォルトポート
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	coteacher.RegisterCoTeacherServiceServer(grpcServer, &services.Server{DB: db}) // サービスの登録

	// サーバの起動
	log.Printf("Starting gRPC server on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

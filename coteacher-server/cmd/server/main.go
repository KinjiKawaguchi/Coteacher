package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/api/grpc_server"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"

	"github.com/go-sql-driver/mysql" //lint:ignore ST1019 this is just example
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/exp/slog"
)

func main() {

	// Open a connection to the database
	db, err := sql.Open("mysql", "gftt3k024uo6qeynrxex:pscale_pw_8UCfiF0ZIc2NLZn9Ap42pNNPkRfF3Dl8ofmLCjd414v@tcp(aws.connect.psdb.cloud)/coteacher?tls=true&interpolateParams=true")
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}

	// PlanetScaleへの接続をテスト
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping PlanetScale: %v", err)
	}
	log.Println("Successfully connected to PlanetScale!")

	// その他の設定
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}
	mysqlConfig := &mysql.Config{
		User:                 "gftt3k024uo6qeynrxex",
		Passwd:               "pscale_pw_8UCfiF0ZIc2NLZn9Ap42pNNPkRfF3Dl8ofmLCjd414v",
		Net:                  "tcp",
		Addr:                 "aws.connect.psdb.cloud",
		DBName:               "coteacher",
		ParseTime:            true,
		Loc:                  jst,
		TLSConfig:            "true", // この行は必要に応じて調整してください。
		AllowNativePasswords: true,
		InterpolateParams:    true,
	}

	log.Println("Starting server...")

	entClient, err := ent.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Running migration...")

	defer entClient.Close()
	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Println("Migration failed!")
		log.Fatal(err)
	}

	log.Println("Migration done!")

	// logger
	logger := slog.Default()

	log.Println("Starting gRPC server...")

	// gRPCサーバの設定
	srv := grpc_server.New(
		grpc_server.WithLogger(logger),
		grpc_server.WithEntClient(entClient),
	)
	lsnr, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	defer lsnr.Close()
	go func() {
		logger.Info("server launched")
		if err := srv.Serve(lsnr); err != nil {
			log.Fatal(err)
		}
	}()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	logger.Info("server is being stopped")
	srv.GracefulStop()
}

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/api/connect_server"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"

	"github.com/go-sql-driver/mysql" //lint:ignore ST1019 this is just example
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/exp/slog"
)

func main() {

	// Open a connection to the database
	db, err := sql.Open("mysql", "hi8tn7ia8rpaej3csgfz:pscale_pw_7OOn6ZqjfHkDQu8NxDXQ0zQQzUg0Fz4QNx4jbTCZ35y@tcp(aws.connect.psdb.cloud)/coteacher?tls=true&interpolateParams=true")
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
		User:                 "hi8tn7ia8rpaej3csgfz",
		Passwd:               "pscale_pw_7OOn6ZqjfHkDQu8NxDXQ0zQQzUg0Fz4QNx4jbTCZ35y",
		Net:                  "tcp",
		Addr:                 "aws.connect.psdb.cloud",
		DBName:               "coteacher",
		ParseTime:            true,
		Loc:                  jst,
		TLSConfig:            "true", // この行は必要に応じて調整してください。
		AllowNativePasswords: true,
		InterpolateParams:    true,
	}

	log.Println("Connecting to PlanetScale...")

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

	log.Println("Starting server...")

	srv := connect_server.New(
		fmt.Sprintf("0.0.0.0:%s", "50051"),
		connect_server.WithLogger(logger),
		connect_server.WithEntClient(entClient),
		connect_server.WithFrontendURL("http://localhost:3000"),
	)

	go func() {
		logger.Info("server launched", slog.String("port", "50051"))
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
	logger.Info("server is being stopped")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal(err)
	}
}

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
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/config"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"

	"github.com/go-sql-driver/mysql" //lint:ignore ST1019 this is just example
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/exp/slog"
)

func main() {
	// コンフィグを読み込む
	config, _ := config.New()

	// MySQLへの接続をテスト
	db, err := sql.Open("mysql", config.DSN)
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}

	// PlanetScaleへの接続をテスト
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping PlanetScale: %v", err)
	}

	// PlanetScaleへの接続を開始
	log.Println("Connecting to PlanetScale...")

	mysqlConfig := &mysql.Config{
		User:                 config.DBUser,
		Passwd:               config.DBPassword,
		Net:                  "tcp",
		Addr:                 config.DBAddress,
		DBName:               config.DBName,
		ParseTime:            true,
		TLSConfig:            "true", // この行は必要に応じて調整してください。
		AllowNativePasswords: true,
		InterpolateParams:    true,
	}

	entClient, err := ent.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully connected to PlanetScale!")

	// マイグレーションを実行
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
		fmt.Sprintf("0.0.0.0:%s", config.Port),
		connect_server.WithLogger(logger),
		connect_server.WithEntClient(entClient),
		connect_server.WithFrontendURL(config.FrontEndURL),
	)

	go func() {
		logger.Info("server launched", slog.String("port", config.Port))
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

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	Config "github.com/KinjiKawaguchi/Coteacher/coteacher-server/config"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/api/connect_server"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"

	"github.com/go-sql-driver/mysql" //lint:ignore ST1019 this is just example
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/exp/slog"
)

func main() {
	config, _ := Config.New()

	// その他の設定
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	mysqlConfig := &mysql.Config{
		User:                 config.DBUser,
		Passwd:               config.DBPassword,
		Net:                  "tcp",
		Addr:                 config.DBAddress,
		DBName:               config.DBName,
		ParseTime:            true,
		Loc:                  jst,
		AllowNativePasswords: true,
		InterpolateParams:    true,
	}

	log.Println("Connecting to CloudSQL...")

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

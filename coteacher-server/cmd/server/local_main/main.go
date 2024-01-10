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
	"github.com/go-sql-driver/mysql" //lint:ignore ST1019 this is just example
	_ "github.com/go-sql-driver/mysql"

	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/api/connect_server"
	"github.com/KinjiKawaguchi/Coteacher/coteacher-server/domain/repository/ent"

	"golang.org/x/exp/slog"
)

func main() {
	config, _ := Config.New()

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}
	mysqlConfig := &mysql.Config{
		DBName:    "coteacher",
		User:      "root",
		Passwd:    "016be2senpu7f",
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
		Loc:       jst,
	}
	entClient, err := ent.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer entClient.Close()
	if err := entClient.Schema.Create(context.Background()); err != nil {
		log.Fatal(err)
	}

	// logger
	logger := slog.Default()

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

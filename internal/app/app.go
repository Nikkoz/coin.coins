package app

import (
	"coins/configs"
	"coins/internal/providers"
	"coins/pkg/types/context"
	"coins/pkg/types/logger"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var config configs.Config

func init() {
	envInit()
	configInit()
}

func Run() {
	ctx := context.Empty()
	defer ctx.Cancel()

	logger.New(config.App.Environment.IsProduction(), config.Log.Level.String())

	// @todo: think about move to service providers
	conn, conClose := connectionDB()
	defer conClose()

	migrate(conn)

	provider := providers.New(ctx, config, conn)
	defer provider.Broker.Close()

	if provider.Grpc == nil {
		return
	}

	var (
		messenger = provider.Delivery.Messenger
		http      = provider.Delivery.Http
		grpc      = provider.Delivery.Grpc
	)

	messenger.Run(provider.Broker, config.Broker.Topics)
	http.Run(config)

	grpcServer, listener := grpc.Run(config)
	if grpcServer == nil {
		return
	}
	defer listener.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err := <-http.Notify():
		logger.Error(fmt.Errorf("app - Run http server: %v", err))
	case err := <-messenger.Notify():
		logger.Error(fmt.Errorf("app - Run msg brocker: %v", err))
	case err := <-grpc.Notify():
		logger.Fatal(fmt.Errorf("app - Run grpc server: %v", err))
	case done := <-ctx.Done():
		logger.Info(fmt.Sprintf("app - Run - ctx.Done: %v", done))
	}
}

func envInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("load environment failed: %v\n", err)
	}
}

func configInit() {
	cfg, err := configs.New()
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	config = *cfg
}

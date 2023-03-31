package app

import (
	"coins/configs"
	messageBroker "coins/internal/delivery/broker"
	"coins/internal/delivery/grpc"
	deliveryHttp "coins/internal/delivery/http"
	repositoryBroker "coins/internal/repository/coin/broker"
	repositoryCoin "coins/internal/repository/coin/database"
	repositoryGrpc "coins/internal/repository/coin/grpc"
	repositoryUrl "coins/internal/repository/url/database"
	coinFactory "coins/internal/useCase/factories/coin"
	urlFactory "coins/internal/useCase/factories/url"
	grpcClient "coins/pkg/grpc"
	"coins/pkg/types/context"
	"coins/pkg/types/logger"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var config *configs.Config

func init() {
	envInit()
	configInit()
}

func Run() {
	ctx := context.Empty()
	defer ctx.Cancel()

	logger.New(config.App.Environment.IsProduction(), config.Log.Level.String())

	conn, conClose := connectionDB()
	defer conClose()

	migrate(conn)

	broker := ConnectionBroker()
	defer broker.Close()

	grpcConn, err := grpcClient.New(ctx, config.Grpc.Host, config.Grpc.Port, config.App.Name, config.App.Version).GetConnection()
	if err != nil {
		_ = logger.ErrorWithContext(ctx, errors.Wrap(err, "failed to create grpc client"))

		return
	}

	var (
		repoUrl      = repositoryUrl.New(conn, repositoryUrl.Options{})
		repoCoin     = repositoryCoin.New(conn, repoUrl, repositoryCoin.Options{})
		repoBroker   = repositoryBroker.New(broker, repositoryBroker.Options{})
		repoGrpc     = repositoryGrpc.New(grpcConn, repositoryGrpc.Options{})
		fCoin        = coinFactory.New(repoCoin, repoBroker, repoGrpc, coinFactory.Options{})
		fUrl         = urlFactory.New(repoUrl, urlFactory.Options{})
		messenger    = messageBroker.New(fCoin, fUrl, messageBroker.Options{})
		listenerHttp = deliveryHttp.New(fCoin, fUrl, deliveryHttp.Options{})
		listenerGrpc = grpc.New(fCoin, grpc.Options{BrokerTopic: config.Broker.Topics[0]})
	)

	messenger.Run(broker, config.Broker.Topics)
	listenerHttp.Run(*config)

	grpcServer, listener := listenerGrpc.Run(config)
	if grpcServer == nil {
		return
	}
	defer listener.Close()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err := <-listenerHttp.Notify():
		logger.Error(fmt.Errorf("app - Run http server: %v", err))
	case err := <-messenger.Notify():
		logger.Error(fmt.Errorf("app - Run msg brocker: %v", err))
	case err := <-listenerGrpc.Notify():
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

	config = cfg
}

package app

import (
	"coins/configs"
	messageBroker "coins/internal/delivery/broker"
	deliveryHttp "coins/internal/delivery/http"
	repositoryBroker "coins/internal/repository/coin/broker"
	repositoryCoin "coins/internal/repository/coin/database"
	repositoryUrl "coins/internal/repository/url/database"
	coinFactory "coins/internal/useCase/factories/coin"
	urlFactory "coins/internal/useCase/factories/url"
	"coins/pkg/types/logger"
	"fmt"
	"github.com/joho/godotenv"
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
	logger.New(*config)

	conn, conClose := connectionDB()
	defer conClose()

	migrate(conn)

	broker := ConnectionBroker()
	defer broker.Close()

	var (
		notify = make(chan error, 1)

		repoUrl      = repositoryUrl.New(conn, repositoryUrl.Options{})
		repoCoin     = repositoryCoin.New(conn, repoUrl, repositoryCoin.Options{})
		repoBroker   = repositoryBroker.New(broker, repositoryBroker.Options{})
		fCoin        = coinFactory.New(repoCoin, repoBroker, coinFactory.Options{})
		fUrl         = urlFactory.New(repoUrl, urlFactory.Options{})
		messenger    = messageBroker.New(fCoin, fUrl, messageBroker.Options{Notify: notify})
		listenerHttp = deliveryHttp.New(fCoin, fUrl, deliveryHttp.Options{Notify: notify})
	)

	messenger.Run(broker, config.Broker.Topics)
	listenerHttp.Run(*config)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err := <-notify:
		logger.Fatal(fmt.Errorf("app - Run: %v", err))
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

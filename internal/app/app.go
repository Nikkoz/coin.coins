package app

import (
	"coins/configs"
	messageBroker "coins/internal/delivery/broker"
	repositoryBroker "coins/internal/repository/coin/broker"
	repositoryCoin "coins/internal/repository/coin/database"
	repositoryUrl "coins/internal/repository/url/database"
	coinFactory "coins/internal/useCase/factories/coin"
	urlFactory "coins/internal/useCase/factories/url"
	"github.com/joho/godotenv"
	"log"
)

var config *configs.Config

func init() {
	envInit()
	configInit()
}

func Run() {
	// @todo: see https://github.com/evrone/go-clean-template/blob/34844d644b3cd20696b7bebbec32b0a65678ba7a/internal/app/app.go
	//log := logger.New(config.Log.Level)

	conn, conClose := connectionDB()
	defer conClose()

	migrate(conn)

	broker := ConnectionBroker()
	defer broker.Close()

	var (
		repoUrl    = repositoryUrl.New(conn, repositoryUrl.Options{})
		repoCoin   = repositoryCoin.New(conn, repoUrl, repositoryCoin.Options{})
		repoBroker = repositoryBroker.New(broker, repositoryBroker.Options{})
		fCoin      = coinFactory.New(repoCoin, repoBroker, coinFactory.Options{})
		fUrl       = urlFactory.New(repoUrl, urlFactory.Options{})
		messenger  = messageBroker.New(fCoin, fUrl, messageBroker.Options{})
	)

	err := messenger.Run(broker, config.Broker.Topics)
	if err != nil {
		log.Fatalf("message broker error: %e", err)
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

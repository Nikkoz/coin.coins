package app

import (
	"coins/configs"
	messageBroker "coins/internal/delivery/broker"
	repositoryBroker "coins/internal/repository/broker"
	repositoryStorage "coins/internal/repository/storage/database"
	coinFactory "coins/internal/useCase/factories/coin"
	urlFactory "coins/internal/useCase/factories/url"
	"log"
)

var config *configs.Config

func init() {
	cfg, err := configs.New()
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	config = cfg
}

func Run() {
	// @todo: see https://github.com/evrone/go-clean-template/blob/34844d644b3cd20696b7bebbec32b0a65678ba7a/internal/app/app.go
	//log := logger.New(config.Log.Level)

	conn, conClose := connectionDB()
	defer conClose()

	broker := ConnectionBroker()
	defer broker.Close()

	var (
		repoStorage = repositoryStorage.New(conn, repositoryStorage.Options{})
		repoBroker  = repositoryBroker.New(broker, repositoryBroker.Options{})
		fCoin       = coinFactory.New(repoStorage, repoBroker, coinFactory.Options{})
		fUrl        = urlFactory.New(repoStorage, urlFactory.Options{})
		messenger   = messageBroker.New(fCoin, fUrl, messageBroker.Options{})
	)

	err := messenger.Run(broker)
	if err != nil {
		log.Fatalf("message broker error: %e", err)
	}

	//Migrate(conn)
}

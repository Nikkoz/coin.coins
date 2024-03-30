package broker

import (
	"coins/configs"
	"coins/pkg/store/messageBroker"
	"log"
)

type ServiceProvider struct{}

func New(config configs.Config) messageBroker.MessageBroker {
	settings := settingsBroker(config)

	return connectionBroker(settings)
}

func connectionBroker(settings messageBroker.Settings) messageBroker.MessageBroker {
	broker, err := messageBroker.New(settings)
	if err != nil {
		log.Fatalf("db error: %v", err)
	}

	return broker
}

func settingsBroker(config configs.Config) messageBroker.Settings {
	return messageBroker.NewSettings(
		config.Broker.Connection,
		config.Broker.Broker,
		config.Broker.SessionTimeout,
		config.Sasl.IsUse,
		config.Sasl.Broker,
		config.Sasl.User,
		config.Sasl.Password,
		config.Sasl.Mechanisms,
		config.Sasl.SecurityProtocol,
		config.SchemaRegistry.Type,
		config.SchemaRegistry.Url,
	)
}

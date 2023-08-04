package app

import (
	"coins/pkg/store/messageBroker"
	"log"
)

func ConnectionBroker() messageBroker.MessageBroker {
	broker, err := messageBroker.New(settingsBroker())
	if err != nil {
		log.Fatalf("db error: %v", err)
	}

	return broker
}

func settingsBroker() messageBroker.Settings {
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

package kafka

import (
	"coins/pkg/store/messageBroker"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func ConfigMap(settings messageBroker.Settings) (*kafka.ConfigMap, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers":  settings.Broker,
		"session.timeout.ms": settings.SessionTimeout,
		"enable.idempotence": true,
		"auto.offset.reset":  "earliest",
		"group.id":           "reader-topics",
		//"client.id":          os.Getenv("CLIENT_ID"),
		"compression.codec": "snappy",
	}

	if settings.Sasl.IsUse {
		err := config.SetKey("bootstrap.servers", settings.Sasl.Broker)
		err = config.SetKey("sasl.username", settings.Sasl.User)
		err = config.SetKey("sasl.password", settings.Sasl.Password)
		err = config.SetKey("sasl.mechanisms", settings.Sasl.Mechanisms)
		err = config.SetKey("security.protocol", settings.Sasl.SecurityProtocol)

		if err != nil {
			return nil, err
		}
	}

	return config, nil
}

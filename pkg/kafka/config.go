package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"os"
)

func generateConfigMap() (*kafka.ConfigMap, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers":  os.Getenv("KAFKA_BROKER"),
		"session.timeout.ms": os.Getenv("KAFKA_SESSION_TIMEOUT"),
		"enable.idempotence": true,
		"auto.offset.reset":  "earliest",
		"group.id":           "reader-topics",
		"client.id":          os.Getenv("CLIENT_ID"),
		"compression.codec":  "snappy",
	}

	if os.Getenv("USE_SASL") == "true" {
		err := config.SetKey("bootstrap.servers", os.Getenv("SASL_BROKER"))
		err = config.SetKey("sasl.username", os.Getenv("SASL_USERNAME"))
		err = config.SetKey("sasl.password", os.Getenv("SASL_PASSWORD"))
		err = config.SetKey("sasl.mechanisms", os.Getenv("SASL_MECHANISMS"))
		err = config.SetKey("security.protocol", os.Getenv("SASL_SECURITY_PROTOCOL"))

		if err != nil {
			return nil, err
		}
	}

	return config, nil
}

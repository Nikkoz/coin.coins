package messageBroker

import "github.com/confluentinc/confluent-kafka-go/kafka"

type (
	Settings struct {
		Connection     string
		Broker         string
		SessionTimeout uint16

		Sasl           Sasl
		SchemaRegistry SchemaRegistry
	}

	Sasl struct {
		IsUse            bool
		Broker           string
		User             string
		Password         string
		Mechanisms       string
		SecurityProtocol string
	}

	SchemaRegistry struct {
		Type string
		Url  string
	}
)

func NewSettings(connection, broker string, timeout uint16, useSasl bool, saslBroker, user, password, mechanisms, protocol, srType, url string) Settings {
	return Settings{
		Connection:     connection,
		Broker:         broker,
		SessionTimeout: timeout,
		Sasl: Sasl{
			IsUse:            useSasl,
			Broker:           saslBroker,
			User:             user,
			Password:         password,
			Mechanisms:       mechanisms,
			SecurityProtocol: protocol,
		},
		SchemaRegistry: SchemaRegistry{
			Type: srType,
			Url:  url,
		},
	}
}

func (settings Settings) ToKafkaConfig() (*kafka.ConfigMap, error) {
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

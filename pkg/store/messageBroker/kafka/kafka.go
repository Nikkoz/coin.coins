package kafka

import (
	"coins/pkg/store/messageBroker"
	"coins/pkg/store/messageBroker/kafka/avro"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"log"
	"os"
)

const (
	AVRO     = "AVRO"
	PROTOBUF = "PROTOBUF"
)

var ErrSchemaNotSupport = errors.New("schema registry type doesn't supported")

type Kafka struct {
	consumer     *kafka.Consumer
	deserializer *serde.Deserializer
}

func New(settings messageBroker.Settings) (messageBroker.MessageBroker, error) {
	config, err := ConfigMap(settings)
	if err != nil {
		return nil, err
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}

	ds, err := deserializer(settings.SchemaRegistry)
	if err != nil {
		return nil, err
	}

	var broker messageBroker.MessageBroker = Kafka{
		consumer:     consumer,
		deserializer: ds,
	}

	return broker, nil
}

func (k Kafka) Subscribe(topics []string) error {
	return k.consumer.SubscribeTopics(topics, nil)
}

func (k Kafka) Consume(sigChan chan os.Signal, callback messageBroker.ConsumeFunc) {
	for {
		select {
		case sig := <-sigChan:
			log.Printf("Caught signal %v: terminating\n", sig)

			break
		default:
			ev := k.consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				err := callback(k.deserializer, e)
				if err != nil {
					log.Printf("%% Error: %v\n", err)
				}
			case kafka.Error:
				log.Printf("%% Error: %v: %v\n", e.Code(), e)
			default:
				log.Printf("Ignored %v\n", e)
			}
		}
	}
}

func (k Kafka) Close() {
	k.consumer.Close()
}

func deserializer(sr messageBroker.SchemaRegistry) (*serde.Deserializer, error) {
	var ds serde.Deserializer
	var err error

	switch sr.Type {
	case AVRO:
		ds, err = avro.Deserializer(sr.Url)
	case PROTOBUF:
		return nil, ErrSchemaNotSupport
	default:
		return nil, ErrSchemaNotSupport
	}

	if err != nil {
		return nil, err
	}

	return &ds, nil
}

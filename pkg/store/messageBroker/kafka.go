package messageBroker

import (
	broker "coins/pkg/store/messageBroker/serde"
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

func NewKafka(settings Settings) (MessageBroker, error) {
	config, err := settings.ToKafkaConfig()
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

	var mb MessageBroker = Kafka{
		consumer:     consumer,
		deserializer: ds,
	}

	return mb, nil
}

func (k Kafka) Subscribe(topics []string) error {
	return k.consumer.SubscribeTopics(topics, nil)
}

func (k Kafka) Consume(sigChan chan os.Signal, doneChan chan bool, callback ConsumeFunc) {
	running := true

	for running {
		select {
		case sig := <-sigChan:
			log.Printf("Caught signal %v: terminating\n", sig)

			running = false
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

				running = false
			default:
				log.Printf("Ignored %v\n", e)
			}
		}
	}

	close(doneChan)
}

func (k Kafka) Close() {
	k.consumer.Close()
}

func deserializer(sr SchemaRegistry) (*serde.Deserializer, error) {
	var ds serde.Deserializer
	var err error

	switch sr.Type {
	case AVRO:
		ds, err = broker.AvroDeserializer(sr.Url)
	case PROTOBUF:
		ds, err = broker.ProtobufDeserializer(sr.Url)
	default:
		return nil, ErrSchemaNotSupport
	}

	if err != nil {
		return nil, err
	}

	return &ds, nil
}

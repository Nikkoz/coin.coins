package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/avro"
	"log"
	"os"
)

type Kafka struct {
	consumer     *kafka.Consumer
	deserializer *avro.SpecificDeserializer
}

type ConsumeFunc func(deserializer *avro.SpecificDeserializer, msg *kafka.Message)

func NewKafka() (*Kafka, error) {
	config, err := generateConfigMap()
	if err != nil {
		return nil, err
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}

	deserializer, err := NewAvroDeserializer()
	if err != nil {
		return nil, err
	}

	return &Kafka{
		consumer:     consumer,
		deserializer: deserializer,
	}, nil
}

func (k *Kafka) Subscribe(topics []string) error {
	return k.consumer.SubscribeTopics(topics, nil)
}

func (k *Kafka) Consume(sigchan chan os.Signal, callback ConsumeFunc) {
	for {
		select {
		case sig := <-sigchan:
			log.Printf("Caught signal %v: terminating\n", sig)

			break
		default:
			ev := k.consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				callback(k.deserializer, e)
			case kafka.Error:
				log.Printf("%% Error: %v: %v\n", e.Code(), e)
			default:
				log.Printf("Ignored %v\n", e)
			}
		}
	}
}

func (k *Kafka) Close() {
	k.consumer.Close()
}

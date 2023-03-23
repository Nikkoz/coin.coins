package messageBroker

import (
	broker "coins/pkg/store/messageBroker/serde"
	"coins/pkg/types/context"
	"coins/pkg/types/logger"
	"errors"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
)

const (
	AVRO     = "AVRO"
	PROTOBUF = "PROTOBUF"
)

var ErrSchemaNotSupport = errors.New("schema registry type doesn't supported")

type Kafka struct {
	consumer     *kafka.Consumer
	deserializer serde.Deserializer
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

func (k Kafka) Consume(notify chan error, c context.Context, callback ConsumeFunc) {
	ctx := c.Copy()
	defer ctx.Cancel()

	for {
		select {
		case <-notify:
			return
		default:
			ev := k.consumer.Poll(100)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				err := callback(ctx, k.deserializer, *e.TopicPartition.Topic, e.Value)
				if err != nil {
					logger.Error(err)
				}
			case kafka.Error:
				notify <- logger.ErrorWithContext(ctx, e)

				return
			default:
				logger.Info(fmt.Sprintf("Ignored %v\n", e))
			}
		}
	}
}

func (k Kafka) Close() {
	err := k.consumer.Close()
	if err != nil {
		return
	}
}

func deserializer(sr SchemaRegistry) (serde.Deserializer, error) {
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

	return ds, nil
}

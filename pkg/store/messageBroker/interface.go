package messageBroker

import (
	"coins/pkg/types/context"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
)

type ConsumeFunc func(ctx context.Context, deserializer serde.Deserializer, topic string, msg []byte) error

type MessageBroker interface {
	Subscribe(topics []string) error
	Consume(doneChan chan error, ctx context.Context, callback ConsumeFunc)
	Close()
}

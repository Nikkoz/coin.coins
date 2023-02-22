package messageBroker

import (
	"coins/pkg/types/context"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"os"
)

type ConsumeFunc func(ctx context.Context, deserializer serde.Deserializer, topic string, msg []byte) error

type MessageBroker interface {
	Subscribe(topics []string) error
	Consume(sigChan chan os.Signal, doneChan chan bool, ctx context.Context, callback ConsumeFunc)
	Close()
}

package messageBroker

import (
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"os"
)

type ConsumeFunc func(deserializer *serde.Deserializer, msg any) error

type MessageBroker interface {
	Subscribe(topics []string) error
	Consume(sigChan chan os.Signal, callback ConsumeFunc)
	Close()
}

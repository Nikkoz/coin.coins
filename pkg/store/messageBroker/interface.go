package messageBroker

import (
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"os"
)

type ConsumeFunc func(deserializer serde.Deserializer, topic string, msg []byte) error

type MessageBroker interface {
	Subscribe(topics []string) error
	Consume(sigChan chan os.Signal, doneChan chan bool, callback ConsumeFunc)
	Close()
}

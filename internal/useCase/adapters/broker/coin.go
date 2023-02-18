package broker

import (
	"coins/internal/domain/coin"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
)

type Coin interface {
	SubscribeCoin(topics []string) error
	ProduceCoin(serializer *serde.Serializer, msg any) error
	ConsumeCoin(deserializer serde.Deserializer, topic string, msg []byte) ([]*coin.Coin, error)
}

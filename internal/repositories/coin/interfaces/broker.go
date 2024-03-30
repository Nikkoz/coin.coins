package interfaces

import (
	"coins/internal/domain/coin"
	"coins/pkg/types/context"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
)

type Broker interface {
	SubscribeCoin(ctx context.Context, topics []string) error
	ProduceCoin(ctx context.Context, serializer *serde.Serializer, msg any) error
	ConsumeCoin(ctx context.Context, deserializer serde.Deserializer, topic string, msg []byte) ([]*coin.Coin, error)
}

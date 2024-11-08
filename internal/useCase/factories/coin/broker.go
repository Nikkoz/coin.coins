package coin

import (
	"coins/pkg/types/context"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
)

func (factory *Factory) Subscribe(ctx context.Context, topics []string) error {
	return factory.adapterBroker.SubscribeCoin(ctx, topics)
}

func (factory *Factory) Consume(ctx context.Context, deserializer serde.Deserializer, topic string, msg []byte) error {
	coins, err := factory.adapterBroker.ConsumeCoin(ctx, deserializer, topic, msg)
	if err != nil {
		return err
	}

	_, err = factory.Upsert(ctx, coins...)

	return err
}

func (factory *Factory) Produce(ctx context.Context, serializer *serde.Serializer, msg any) error {
	return factory.adapterBroker.ProduceCoin(ctx, serializer, msg)
}

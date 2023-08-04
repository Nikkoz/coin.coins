package broker

import (
	"coins/internal/domain/coin"
	"coins/internal/repository/coin/broker/entities"
	"coins/pkg/types/context"
	"coins/pkg/types/logger"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
)

func (r *Repository) SubscribeCoin(ctx context.Context, topics []string) error {
	defer ctx.Copy().Cancel()

	err := r.broker.Subscribe(topics)
	if err != nil {
		return logger.FatalWithContext(ctx, err)
	}

	return nil
}

func (r *Repository) ProduceCoin(ctx context.Context, serializer *serde.Serializer, msg any) error {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) ConsumeCoin(ctx context.Context, deserializer serde.Deserializer, topic string, msg []byte) ([]*coin.Coin, error) {
	defer ctx.Copy().Cancel()

	value := entities.NewCoins()
	if err := deserializer.DeserializeInto(topic, msg, &value); err != nil {
		return nil, logger.ErrorWithContext(ctx, err)
	}

	//if e.Headers != nil {
	//	fmt.Printf("%% Headers: %v\n", e.Headers)
	//}

	coins, err := toModels(value.Coins)
	if err != nil {
		return nil, logger.FatalWithContext(ctx, err)
	}

	return coins, nil
}

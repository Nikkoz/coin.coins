package broker

import (
	"coins/internal/domain/coin"
	"coins/internal/repository/coin/broker/entities"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
)

func (r *Repository) SubscribeCoin(topics []string) error {
	return r.broker.Subscribe(topics)
}

func (r *Repository) ProduceCoin(serializer *serde.Serializer, msg any) error {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) ConsumeCoin(deserializer serde.Deserializer, topic string, msg []byte) ([]*coin.Coin, error) {
	value := entities.NewCoins()
	if err := deserializer.DeserializeInto(topic, msg, &value); err != nil {
		return nil, fmt.Errorf("failed to deserialize payload: %s\n", err)
	}

	//if e.Headers != nil {
	//	fmt.Printf("%% Headers: %v\n", e.Headers)
	//}

	coins, err := toModels(value.Coins)
	if err != nil {
		return nil, fmt.Errorf("validation error: %v\n", err)
	}

	return coins, nil
}

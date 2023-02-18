package coin

import "github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"

func (factory *Factory) Subscribe(topics []string) error {
	return factory.adapterBroker.SubscribeCoin(topics)
}

func (factory *Factory) Consume(deserializer serde.Deserializer, topic string, msg []byte) error {
	coins, err := factory.adapterBroker.ConsumeCoin(deserializer, topic, msg)
	if err != nil {
		return err
	}

	return factory.Upsert(coins...)
}

func (factory *Factory) Produce(serializer *serde.Serializer, msg any) error {
	return factory.adapterBroker.ProduceCoin(serializer, msg)
}

package broker

import "github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"

type Coin interface {
	Subscribe(topics []string) error
	Produce(serializer *serde.Serializer, msg any) error
	Consume(deserializer *serde.Deserializer, msg any) error
}

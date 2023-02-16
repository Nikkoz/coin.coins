package broker

import "github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"

func (r *Repository) Subscribe(topics []string) error {
	return r.broker.Subscribe(topics)
}

func (r *Repository) Produce(serializer *serde.Serializer, msg any) error {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) Consume(deserializer *serde.Deserializer, msg any) error {
	// TODO implement me
	panic("implement me")
}

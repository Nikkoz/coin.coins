package serde

import (
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/protobuf"
	"log"
)

func ProtobufDeserializer(url string) (*protobuf.Deserializer, error) {
	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(url))
	if err != nil {
		log.Printf("Failed to create schema registry client: %v", err)

		return nil, err
	}

	return protobuf.NewDeserializer(client, serde.ValueSerde, protobuf.NewDeserializerConfig())
}

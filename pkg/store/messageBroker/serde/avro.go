package serde

import (
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/avro"
	"log"
)

func AvroDeserializer(url string) (*avro.SpecificDeserializer, error) {
	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(url))
	if err != nil {
		log.Printf("Failed to create schema registry client: %v", err)

		return nil, err
	}

	return avro.NewSpecificDeserializer(client, serde.ValueSerde, avro.NewDeserializerConfig())
}

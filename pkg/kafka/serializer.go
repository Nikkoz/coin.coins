package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/avro"
	"log"
	"os"
)

func NewAvroSerializer() (*avro.SpecificSerializer, error) {
	if os.Getenv("USE_AVRO") == "false" {
		return nil, nil
	}

	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(os.Getenv("SCHEMA_REGISTRY_URL")))
	if err != nil {
		log.Fatalf("Failed to create schema registry client: %v", err)

		return nil, err
	}

	return avro.NewSpecificSerializer(client, serde.ValueSerde, avro.NewSerializerConfig())
}

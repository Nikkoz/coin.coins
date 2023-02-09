package handlers

import (
	"coins/internals/entities"
	"coins/internals/models"
	"coins/internals/models/factories"
	"coins/pkg/db"
	"coins/pkg/db/helpers/pg"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/avro"
	"gorm.io/gorm"
	"log"
)

func Consume(deserializer *avro.SpecificDeserializer, msg *kafka.Message) {
	value := entities.NewCoins()
	if err := deserializer.DeserializeInto(*msg.TopicPartition.Topic, msg.Value, &value); err != nil {
		log.Printf("Failed to deserialize payload: %s\n", err)

		return
	}

	//if e.Headers != nil {
	//	fmt.Printf("%% Headers: %v\n", e.Headers)
	//}

	coins := prepareCoins(value.Coins)
	if err := factories.NewCoinFactory(dbConn()).Upsert(coins); err != nil {
		log.Printf("Failed to deserialize payload: %s\n", err)
	}
}

func prepareCoins(coins []entities.Coin) []models.Coin {
	newCoins := make([]models.Coin, 0, len(coins))

	for _, data := range coins {
		coin := models.Coin{
			ID:   uint(data.Id),
			Name: data.Name,
			Code: data.Code,
			Icon: data.Icon,
		}

		var urls []*models.CoinUrl
		if len(data.Urls) > 0 {
			for _, u := range data.Urls {
				url := &models.CoinUrl{
					ExternalID: uint(u.Id),
					Link:       u.Link,
					Type:       u.Type,
				}

				urls = append(urls, url)
			}
		}

		coin.CoinUrls = urls
		newCoins = append(newCoins, coin)
	}

	return newCoins
}

func dbConn() *gorm.DB {
	config := db.NewConfig()

	return pg.Open(config)
}

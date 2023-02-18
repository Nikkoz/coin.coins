package interfaces

import (
	"coins/internal/domain/coin"
	"coins/pkg/types/queryParameter"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
)

type (
	Coin interface {
		Save(coin *coin.Coin) (*coin.Coin, error)
		Delete(ID uint) error
		Upsert(coins ...*coin.Coin) error

		CoinReader
		BrokerCoin
	}

	CoinReader interface {
		List(parameter queryParameter.QueryParameter) ([]*coin.Coin, error)
		Count( /*Тут можно передавать фильтр*/ ) (uint64, error)
	}

	BrokerCoin interface {
		Subscribe(topics []string) error
		Consume(deserializer serde.Deserializer, topic string, msg []byte) error
		Produce(serializer *serde.Serializer, msg any) error
	}
)

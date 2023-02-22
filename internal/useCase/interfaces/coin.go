package interfaces

import (
	"coins/internal/domain/coin"
	"coins/pkg/types/context"
	"coins/pkg/types/queryParameter"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
)

type (
	Coin interface {
		Create(ctx context.Context, coin *coin.Coin) (*coin.Coin, error)
		Update(ctx context.Context, coin *coin.Coin) (*coin.Coin, error)
		Delete(ctx context.Context, ID uint) error
		Upsert(ctx context.Context, coins ...*coin.Coin) error

		CoinReader
		BrokerCoin
	}

	CoinReader interface {
		List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*coin.Coin, error)
		Count(ctx context.Context /*Тут можно передавать фильтр*/) (uint64, error)
	}

	BrokerCoin interface {
		Subscribe(ctx context.Context, topics []string) error
		Consume(ctx context.Context, deserializer serde.Deserializer, topic string, msg []byte) error
		Produce(ctx context.Context, serializer *serde.Serializer, msg any) error
	}
)

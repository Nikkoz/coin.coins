package coin

import (
	"coins/internal/domain/coin"
	"coins/pkg/types/queryParameter"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
)

func (factory *Factory) Save(coin *coin.Coin) (*coin.Coin, error) {
	// TODO implement me
	panic("implement me")
}

func (factory *Factory) Delete(ID uint) error {
	// TODO implement me
	panic("implement me")
}

func (factory *Factory) Upsert(coins ...*coin.Coin) ([]*coin.Coin, error) {
	// TODO implement me
	panic("implement me")
}

func (factory *Factory) List(parameter queryParameter.QueryParameter) ([]*coin.Coin, error) {
	// TODO implement me
	panic("implement me")
}

func (factory *Factory) Count( /*Тут можно передавать фильтр*/ ) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

// @todo: возможно нужно разделить на подразделы(db и broker)
func (factory *Factory) Subscribe(topics []string) error {
	return factory.adapterBroker.Subscribe(topics)
}

func (factory *Factory) Consume(deserializer *serde.Deserializer, msg any) error {
	// TODO implement me
	panic("implement me")
}

func (factory *Factory) Produce() error {
	// TODO implement me
	panic("implement me")
}

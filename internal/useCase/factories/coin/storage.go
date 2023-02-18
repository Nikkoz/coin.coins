package coin

import (
	"coins/internal/domain/coin"
	"coins/pkg/types/queryParameter"
)

func (factory *Factory) Create(coin *coin.Coin) (*coin.Coin, error) {
	return factory.adapterStorage.CreateCoin(coin)
}

func (factory *Factory) Update(coin *coin.Coin) (*coin.Coin, error) {
	return factory.adapterStorage.UpdateCoin(coin)
}

func (factory *Factory) Delete(ID uint) error {
	return factory.adapterStorage.DeleteCoin(ID)
}

func (factory *Factory) Upsert(coins ...*coin.Coin) error {
	return factory.adapterStorage.UpsertCoins(coins...)
}

func (factory *Factory) List(parameter queryParameter.QueryParameter) ([]*coin.Coin, error) {
	// TODO implement me
	panic("implement me")
}

func (factory *Factory) Count( /*Тут можно передавать фильтр*/ ) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

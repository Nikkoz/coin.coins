package coin

import (
	"coins/internal/domain/coin"
	"coins/pkg/types/context"
	"coins/pkg/types/queryParameter"
)

func (factory *Factory) Create(ctx context.Context, coin *coin.Coin) (*coin.Coin, error) {
	return factory.adapterStorage.CreateCoin(ctx, coin)
}

func (factory *Factory) Update(ctx context.Context, coin *coin.Coin) (*coin.Coin, error) {
	return factory.adapterStorage.UpdateCoin(ctx, coin)
}

func (factory *Factory) Delete(ctx context.Context, ID uint) error {
	return factory.adapterStorage.DeleteCoin(ctx, ID)
}

func (factory *Factory) Upsert(ctx context.Context, coins ...*coin.Coin) ([]*coin.Coin, error) {
	return factory.adapterStorage.UpsertCoins(ctx, coins...)
}

func (factory *Factory) List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*coin.Coin, error) {
	return factory.adapterStorage.ListCoins(ctx, parameter)
}

func (factory *Factory) Count(ctx context.Context /*Тут можно передавать фильтр*/) (uint64, error) {
	return factory.adapterStorage.CountCoins(ctx)
}

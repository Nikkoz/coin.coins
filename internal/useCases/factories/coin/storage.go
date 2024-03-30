package coin

import (
	"coins/internal/domain/coin"
	"coins/pkg/types/context"
	"coins/pkg/types/queryParameter"
	"github.com/Nikkoz/coin.sync/pkg/protobuf/coins"
)

func (factory *Factory) Create(ctx context.Context, coin *coin.Coin) (*coin.Coin, error) {
	return factory.storage.CreateCoin(ctx, coin)
}

func (factory *Factory) Update(ctx context.Context, coin *coin.Coin) (*coin.Coin, error) {
	return factory.storage.UpdateCoin(ctx, coin)
}

func (factory *Factory) Delete(ctx context.Context, ID uint) error {
	return factory.storage.DeleteCoin(ctx, ID)
}

func (factory *Factory) Upsert(ctx context.Context, coins ...*coin.Coin) ([]*coin.Coin, error) {
	return factory.storage.UpsertCoins(ctx, coins...)
}

func (factory *Factory) FindByID(ctx context.Context, ID uint) (*coin.Coin, error) {
	return factory.storage.CoinByID(ctx, ID)
}

func (factory *Factory) FindFullByID(ctx context.Context, ID uint) (*coin.Coin, *coins.Coin, error) {
	model, err := factory.FindByID(ctx, ID)
	if err != nil {
		return nil, nil, err
	}

	c, err := factory.grpc.GetInfo(ctx, uint64(ID))
	if err != nil {
		return model, nil, nil
	}

	return model, c, nil
}

func (factory *Factory) List(ctx context.Context, parameter queryParameter.QueryParameter) ([]*coin.Coin, error) {
	return factory.storage.ListCoins(ctx, parameter)
}

func (factory *Factory) Count(ctx context.Context /*Тут можно передавать фильтр*/) (uint64, error) {
	return factory.storage.CountCoins(ctx)
}

package storage

import (
	"coins/internal/domain/coin"
	"coins/pkg/types/context"
	"coins/pkg/types/queryParameter"
)

type (
	Coin interface {
		CreateCoin(ctx context.Context, coin *coin.Coin) (*coin.Coin, error)
		UpdateCoin(ctx context.Context, coin *coin.Coin) (*coin.Coin, error)
		DeleteCoin(ctx context.Context, ID uint) error
		UpsertCoins(ctx context.Context, coins ...*coin.Coin) ([]*coin.Coin, error)

		CoinReader
	}

	CoinReader interface {
		CoinByID(ctx context.Context, id uint) (*coin.Coin, error)
		ListCoins(ctx context.Context, parameter queryParameter.QueryParameter) ([]*coin.Coin, error)
		CountCoins(ctx context.Context /*Тут можно передавать фильтр*/) (uint64, error)
	}
)

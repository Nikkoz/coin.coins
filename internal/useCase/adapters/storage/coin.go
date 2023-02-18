package storage

import (
	"coins/internal/domain/coin"
	"coins/pkg/types/queryParameter"
)

type (
	Coin interface {
		CreateCoin(coin *coin.Coin) (*coin.Coin, error)
		UpdateCoin(coin *coin.Coin) (*coin.Coin, error)
		DeleteCoin(ID uint) error
		UpsertCoins(coins ...*coin.Coin) error

		CoinReader
	}

	CoinReader interface {
		ListCoins(parameter queryParameter.QueryParameter) ([]*coin.Coin, error)
		CountCoins( /*Тут можно передавать фильтр*/ ) (uint64, error)
	}
)

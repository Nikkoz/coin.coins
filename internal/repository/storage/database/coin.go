package database

import (
	"coins/internal/domain/coin"
	"coins/pkg/types/queryParameter"
)

func (r *Repository) CreateCoin(coin *coin.Coin) (*coin.Coin, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) UpdateCoin(coin *coin.Coin) (*coin.Coin, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) DeleteCoin(ID uint) error {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) UpsertCoins(coins ...*coin.Coin) ([]*coin.Coin, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) ListCoins(parameter queryParameter.QueryParameter) ([]*coin.Coin, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) CountCoins( /*Тут можно передавать фильтр*/ ) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

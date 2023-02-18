package database

import (
	"coins/internal/domain/coin"
	"coins/internal/domain/url"
	"coins/pkg/types/queryParameter"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (r *Repository) UpsertCoins(coins ...*coin.Coin) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		err := r.db.Omit(clause.Associations).
			Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "id"}},
				UpdateAll: true,
			}).
			Create(coins).
			Error

		if err != nil {
			return err
		}

		return r.saveAssociations(coins...)
	})

	return err
}

func (r *Repository) ListCoins(parameter queryParameter.QueryParameter) ([]*coin.Coin, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) CountCoins( /*Тут можно передавать фильтр*/ ) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) saveAssociations(coins ...*coin.Coin) error {
	urls := make([]*url.Url, 0)

	for _, c := range coins {
		if len(c.CoinUrls) == 0 {
			return nil
		}

		for _, link := range c.CoinUrls {
			link.CoinID = c.ID

			urls = append(urls, link)
		}
	}

	if err := r.repoUrl.UpsertUrls(urls...); err != nil {
		return err
	}

	return nil
}

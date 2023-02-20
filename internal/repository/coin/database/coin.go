package database

import (
	"coins/internal/domain/coin"
	"coins/internal/domain/url"
	"coins/pkg/types/columnCode"
	"coins/pkg/types/queryParameter"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var mappingSort = map[columnCode.ColumnCode]string{
	"id":   "id",
	"name": "name",
	"code": "code",
}

func (r *Repository) CreateCoin(coin *coin.Coin) (*coin.Coin, error) {
	if err := r.db.Create(&coin).Error; err != nil {
		return nil, err
	}

	return coin, nil
}

func (r *Repository) UpdateCoin(coin *coin.Coin) (*coin.Coin, error) {
	if err := r.db.Model(&coin).Save(&coin).Error; err != nil {
		return nil, err
	}

	return coin, nil
}

func (r *Repository) DeleteCoin(ID uint) error {
	return r.db.Delete(&coin.Coin{}, ID).Error
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
	var coins []*coin.Coin

	builder := r.db.Model(&coins)

	if len(parameter.Sorts) > 0 {
		for _, value := range parameter.Sorts.Parsing(mappingSort) {
			if value == "" {
				continue
			}

			builder = builder.Order(value)
		}
	}

	if parameter.Pagination.Limit > 0 {
		builder = builder.Limit(int(parameter.Pagination.Limit))
	}

	if parameter.Pagination.Offset > 0 {
		builder = builder.Offset(int(parameter.Pagination.Offset))
	}

	result := builder.Find(&coins)

	return coins, result.Error
}

func (r *Repository) CountCoins( /*Тут можно передавать фильтр*/ ) (uint64, error) {
	var count int64

	result := r.db.Model(&coin.Coin{}).Count(&count)

	return uint64(count), result.Error
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

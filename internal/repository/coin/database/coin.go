package database

import (
	"coins/internal/domain/coin"
	"coins/internal/domain/url"
	"coins/pkg/store/db/scoupes"
	"coins/pkg/types/columnCode"
	"coins/pkg/types/context"
	"coins/pkg/types/logger"
	"coins/pkg/types/queryParameter"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var mappingSort = map[columnCode.ColumnCode]string{
	"id":   "id",
	"name": "name",
	"code": "code",
}

func (r *Repository) CreateCoin(c context.Context, coin *coin.Coin) (*coin.Coin, error) {
	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	if err := r.db.Create(&coin).Error; err != nil {
		return nil, logger.ErrorWithContext(ctx, err)
	}

	return coin, nil
}

func (r *Repository) UpdateCoin(c context.Context, coin *coin.Coin) (*coin.Coin, error) {
	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	if err := r.db.Model(&coin).Save(&coin).Error; err != nil {
		return nil, logger.ErrorWithContext(ctx, err)
	}

	return coin, nil
}

func (r *Repository) DeleteCoin(c context.Context, ID uint) error {
	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	result := r.db.Delete(&coin.Coin{}, ID)
	if result.Error != nil {
		return logger.ErrorWithContext(ctx, result.Error)
	}

	return nil
}

func (r *Repository) UpsertCoins(c context.Context, coins ...*coin.Coin) error {
	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	err := r.db.Transaction(func(tx *gorm.DB) error {
		err := r.db.Omit(clause.Associations).
			Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "id"}},
				UpdateAll: true,
			}).
			Create(coins).
			Error

		if err != nil {
			return logger.ErrorWithContext(ctx, err)
		}

		return r.saveAssociations(ctx, coins...)
	})

	return logger.ErrorWithContext(ctx, err)
}

func (r *Repository) ListCoins(c context.Context, parameter queryParameter.QueryParameter) ([]*coin.Coin, error) {
	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

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

	if parameter.Pagination.Limit == 0 {
		parameter.Pagination.Limit = r.options.DefaultLimit
	}

	result := builder.
		Scopes(scoupes.Paginate(
			parameter.Pagination.Limit,
			parameter.Pagination.Page,
		)).
		Find(&coins)

	if result.Error != nil {
		return coins, logger.FatalWithContext(ctx, result.Error)
	}

	return coins, nil
}

func (r *Repository) CountCoins(c context.Context /*Тут можно передавать фильтр*/) (uint64, error) {
	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	var count int64

	result := r.db.Model(&coin.Coin{}).Count(&count)

	if result.Error != nil {
		return 0, logger.FatalWithContext(ctx, result.Error)
	}

	return uint64(count), nil
}

func (r *Repository) saveAssociations(c context.Context, coins ...*coin.Coin) error {
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

	if err := r.repoUrl.UpsertUrls(c, urls...); err != nil {
		return logger.ErrorWithContext(c, err)
	}

	return nil
}

package database

import (
	domain "coins/internal/domain/url"
	"coins/pkg/store/db/scoupes"
	"coins/pkg/types/columnCode"
	"coins/pkg/types/queryParameter"
	"gorm.io/gorm/clause"
)

var mappingSort = map[columnCode.ColumnCode]string{
	"id":   "id",
	"type": "type",
}

func (r *Repository) CreateUrl(url *domain.Url) (*domain.Url, error) {
	if err := r.db.Create(&url).Error; err != nil {
		return nil, err
	}

	return url, nil
}

func (r *Repository) UpdateUrl(url *domain.Url) (*domain.Url, error) {
	if err := r.db.Model(&url).Save(&url).Error; err != nil {
		return nil, err
	}

	return url, nil
}

func (r *Repository) DeleteUrl(ID uint) error {
	return r.db.Delete(&domain.Url{}, ID).Error
}

func (r *Repository) UpsertUrls(urls ...*domain.Url) error {
	return r.db.
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "external_id"}},
			UpdateAll: true,
		}).
		Create(urls).
		Error
}

func (r *Repository) UrlById(ID uint) (*domain.Url, error) {
	var url *domain.Url

	result := r.db.First(&url, ID)

	return url, result.Error
}

func (r *Repository) ListUrls(coinId uint, parameter queryParameter.QueryParameter) ([]*domain.Url, error) {
	var urls []*domain.Url

	builder := r.db.Model(&urls)

	if len(parameter.Sorts) > 0 {
		for _, value := range parameter.Sorts.Parsing(mappingSort) {
			if value == "" {
				continue
			}

			builder = builder.Order(value)
		}
	}

	result := builder.
		Where(&domain.Url{CoinID: coinId}).
		Scopes(scoupes.Paginate(
			parameter.Pagination.Limit,
			parameter.Pagination.Page,
		)).
		Find(&urls)

	return urls, result.Error
}

func (r *Repository) CountUrls(coinId uint /*Тут можно передавать фильтр*/) (uint64, error) {
	var count int64
	url := domain.Url{CoinID: coinId}

	result := r.db.Model(&url).Where(&url).Count(&count)

	return uint64(count), result.Error
}

package database

import (
	domain "coins/internal/domain/url"
	"coins/pkg/types/queryParameter"
	"gorm.io/gorm/clause"
)

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

func (r *Repository) ListUrls(parameter queryParameter.QueryParameter) ([]*domain.Url, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) CountUrls( /*Тут можно передавать фильтр*/ ) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

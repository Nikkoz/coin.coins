package database

import (
	"coins/internal/domain/url"
	"coins/pkg/types/queryParameter"
)

func (r *Repository) CreateUrl(url *url.Url) (*url.Url, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) UpdateUrl(url *url.Url) (*url.Url, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) DeleteUrl(ID uint) error {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) UpsertUrls(urls ...*url.Url) ([]*url.Url, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) ListUrls(parameter queryParameter.QueryParameter) ([]*url.Url, error) {
	// TODO implement me
	panic("implement me")
}

func (r *Repository) CountUrls( /*Тут можно передавать фильтр*/ ) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

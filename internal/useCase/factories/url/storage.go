package url

import (
	domain "coins/internal/domain/url"
	"coins/pkg/types/queryParameter"
	"github.com/pkg/errors"
)

func (f *Factory) Create(url *domain.Url) (*domain.Url, error) {
	return f.adapterStorage.CreateUrl(url)
}

func (f *Factory) Update(url *domain.Url) (*domain.Url, error) {
	return f.adapterStorage.UpdateUrl(url)
}

func (f *Factory) Delete(ID uint, coinId uint) error {
	url, err := f.ById(ID)
	if err != nil {
		return err
	}

	if url.CoinID != coinId {
		return errors.New("url doesn't belong to the specified coin")
	}

	return f.adapterStorage.DeleteUrl(ID)
}

func (f *Factory) Upsert(urls ...*domain.Url) error {
	return f.adapterStorage.UpsertUrls(urls...)
}

func (f *Factory) ById(ID uint) (*domain.Url, error) {
	return f.adapterStorage.UrlById(ID)
}

func (f *Factory) List(coinId uint, parameter queryParameter.QueryParameter) ([]*domain.Url, error) {
	return f.adapterStorage.ListUrls(coinId, parameter)
}

func (f *Factory) Count(coinId uint /*Тут можно передавать фильтр*/) (uint64, error) {
	return f.adapterStorage.CountUrls(coinId)
}

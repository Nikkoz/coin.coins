package url

import (
	"coins/internal/domain/url"
	"coins/pkg/types/queryParameter"
)

func (f *Factory) Create(url *url.Url) (*url.Url, error) {
	return f.adapterStorage.CreateUrl(url)
}

func (f *Factory) Update(url *url.Url) (*url.Url, error) {
	// TODO implement me
	panic("implement me")
}

func (f *Factory) Delete(ID uint) error {
	// TODO implement me
	panic("implement me")
}

func (f *Factory) Upsert(urls ...*url.Url) error {
	// TODO implement me
	panic("implement me")
}

func (f *Factory) List(parameter queryParameter.QueryParameter) ([]*url.Url, error) {
	// TODO implement me
	panic("implement me")
}

func (f *Factory) Count( /*Тут можно передавать фильтр*/ ) (uint64, error) {
	// TODO implement me
	panic("implement me")
}

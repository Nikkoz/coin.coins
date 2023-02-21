package interfaces

import (
	"coins/internal/domain/url"
	"coins/pkg/types/queryParameter"
)

type (
	Url interface {
		Create(url *url.Url) (*url.Url, error)
		Update(url *url.Url) (*url.Url, error)
		Delete(ID uint, coinId uint) error
		Upsert(urls ...*url.Url) error

		UrlReader
	}

	UrlReader interface {
		ById(ID uint) (*url.Url, error)
		List(parameter queryParameter.QueryParameter) ([]*url.Url, error)
		Count( /*Тут можно передавать фильтр*/ ) (uint64, error)
	}
)

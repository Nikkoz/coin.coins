package interfaces

import (
	"coins/internal/domain/url"
	"coins/pkg/types/queryParameter"
)

type (
	Url interface {
		Save(url *url.Url) (*url.Url, error)
		Delete(ID uint) error
		Upsert(urls ...*url.Url) ([]*url.Url, error)

		UrlReader
	}

	UrlReader interface {
		List(parameter queryParameter.QueryParameter) ([]*url.Url, error)
		Count( /*Тут можно передавать фильтр*/ ) (uint64, error)
	}
)

package interfaces

import (
	"coins/internal/domain/url"
	"coins/pkg/types/context"
	"coins/pkg/types/queryParameter"
)

type (
	Url interface {
		Create(ctx context.Context, url *url.Url) (*url.Url, error)
		Update(ctx context.Context, url *url.Url) (*url.Url, error)
		Delete(ctx context.Context, ID uint, coinId uint) error
		Upsert(ctx context.Context, urls ...*url.Url) ([]*url.Url, error)

		UrlReader
	}

	UrlReader interface {
		ById(ctx context.Context, ID uint) (*url.Url, error)
		List(ctx context.Context, coinId uint, parameter queryParameter.QueryParameter) ([]*url.Url, error)
		Count(ctx context.Context, coinId uint /*Тут можно передавать фильтр*/) (uint64, error)
	}
)

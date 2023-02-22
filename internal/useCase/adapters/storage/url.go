package storage

import (
	"coins/internal/domain/url"
	"coins/pkg/types/context"
	"coins/pkg/types/queryParameter"
)

type (
	Url interface {
		CreateUrl(ctx context.Context, url *url.Url) (*url.Url, error)
		UpdateUrl(ctx context.Context, url *url.Url) (*url.Url, error)
		DeleteUrl(ctx context.Context, ID uint) error
		UpsertUrls(ctx context.Context, urls ...*url.Url) error

		UrlReader
	}

	UrlReader interface {
		UrlById(ctx context.Context, ID uint) (*url.Url, error)
		ListUrls(ctx context.Context, coinId uint, parameter queryParameter.QueryParameter) ([]*url.Url, error)
		CountUrls(ctx context.Context, coinId uint /*Тут можно передавать фильтр*/) (uint64, error)
	}
)

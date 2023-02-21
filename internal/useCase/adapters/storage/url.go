package storage

import (
	"coins/internal/domain/url"
	"coins/pkg/types/queryParameter"
)

type (
	Url interface {
		CreateUrl(url *url.Url) (*url.Url, error)
		UpdateUrl(url *url.Url) (*url.Url, error)
		DeleteUrl(ID uint) error
		UpsertUrls(urls ...*url.Url) error

		UrlReader
	}

	UrlReader interface {
		UrlById(ID uint) (*url.Url, error)
		ListUrls(coinId uint, parameter queryParameter.QueryParameter) ([]*url.Url, error)
		CountUrls(coinId uint /*Тут можно передавать фильтр*/) (uint64, error)
	}
)

package url

import (
	domain "coins/internal/domain/url"
	"coins/pkg/types/context"
	"coins/pkg/types/logger"
	"coins/pkg/types/queryParameter"
	"github.com/pkg/errors"
)

func (f *Factory) Create(ctx context.Context, url *domain.Url) (*domain.Url, error) {
	return f.storage.CreateUrl(ctx, url)
}

func (f *Factory) Update(ctx context.Context, url *domain.Url) (*domain.Url, error) {
	return f.storage.UpdateUrl(ctx, url)
}

func (f *Factory) Delete(ctx context.Context, ID uint, coinId uint) error {
	url, err := f.ById(ctx, ID)
	if err != nil {
		return err
	}

	if url.CoinID != coinId {
		return logger.ErrorWithContext(ctx, errors.New("url doesn't belong to the specified coin"))
	}

	return f.storage.DeleteUrl(ctx, ID)
}

func (f *Factory) Upsert(ctx context.Context, urls ...*domain.Url) ([]*domain.Url, error) {
	return f.storage.UpsertUrls(ctx, urls...)
}

func (f *Factory) ById(ctx context.Context, ID uint) (*domain.Url, error) {
	return f.storage.UrlById(ctx, ID)
}

func (f *Factory) List(ctx context.Context, coinId uint, parameter queryParameter.QueryParameter) ([]*domain.Url, error) {
	return f.storage.ListUrls(ctx, coinId, parameter)
}

func (f *Factory) Count(ctx context.Context, coinId uint /*Тут можно передавать фильтр*/) (uint64, error) {
	return f.storage.CountUrls(ctx, coinId)
}

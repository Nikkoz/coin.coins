package url

import (
	domain "coins/internal/domain/url"
	"coins/internal/domain/url/types/externalId"
	"coins/internal/domain/url/types/link"
	"coins/internal/domain/url/types/socialMedia"
	mockStorage "coins/internal/repository/url/database/mock"
	"coins/pkg/types/context"
	"coins/pkg/types/queryParameter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var (
	mainUrl    *domain.Url
	repository = new(mockStorage.Url)
	factory    *Factory
)

func TestCreate(t *testing.T) {
	assertion := arrangeCreate(t)

	t.Run("Create url", func(t *testing.T) {
		ctx := context.Empty()

		result, err := factory.Create(ctx, mainUrl)

		assertion.NoError(err)
		assertion.Equal(mainUrl.SocialMedia, result.SocialMedia)
		assertion.Equal(mainUrl.Link, result.Link)
		assertion.Equal(mainUrl.ExternalID, result.ExternalID)
		assertion.NotZero(result, "ID")
	})
}

func TestUpdate(t *testing.T) {
	assertion := arrangeUpdate(t)

	t.Run("Update url", func(t *testing.T) {
		ctx := context.Empty()

		result, err := factory.Update(ctx, mainUrl)

		assertion.NoError(err)
		assertion.Equal(mainUrl, result)
	})
}

func TestDelete(t *testing.T) {
	assertion := arrangeDelete(t)

	t.Run("Delete url", func(t *testing.T) {
		ctx := context.Empty()

		err := factory.Delete(ctx, mainUrl.ID, mainUrl.CoinID)
		assertion.NoError(err)
	})
}

func TestUpsertCreate(t *testing.T) {
	assertion := arrangeUpsertCreate(t)

	t.Run("Upsert(create) url", func(t *testing.T) {
		ctx := context.Empty()

		c := *mainUrl
		result, err := factory.Upsert(ctx, &c)

		assertion.NoError(err)
		assertion.Equal(mainUrl.SocialMedia, result[0].SocialMedia)
		assertion.Equal(mainUrl.Link, result[0].Link)
		assertion.Equal(mainUrl.ExternalID, result[0].ExternalID)
		assertion.NotZero(result[0], "ID")
	})
}

func TestUpsertUpdate(t *testing.T) {
	assertion := arrangeUpsertUpdate(t)

	t.Run("Upsert(create) url", func(t *testing.T) {
		ctx := context.Empty()

		c := *mainUrl
		result, err := factory.Upsert(ctx, &c)

		assertion.NoError(err)
		assertion.Equal(mainUrl, result[0])
	})
}

func TestList(t *testing.T) {
	assertion := arrangeList(t)

	t.Run("List urls", func(t *testing.T) {
		ctx := context.Empty()

		result, err := factory.List(ctx, mainUrl.CoinID, queryParameter.QueryParameter{})

		assertion.NoError(err)
		assertion.Equal(mainUrl, result[0])
	})
}

func TestCount(t *testing.T) {
	assertion := arrangeCount(t)

	t.Run("Count urls", func(t *testing.T) {
		ctx := context.Empty()

		result, err := factory.Count(ctx, mainUrl.CoinID)

		assertion.NoError(err)
		assertion.Equal(uint64(1), result)
	})
}

func arrangeCreate(t *testing.T) *assert.Assertions {
	createUrl(nil)

	factory = New(repository, Options{})
	assertion := assert.New(t)

	repository.
		On(
			"CreateUrl",
			mock.Anything,
			mock.AnythingOfType("*url.Url"),
		).
		Return(func(ctx context.Context, url *domain.Url) *domain.Url {
			assertion.Equal(mainUrl, url)

			url.ID = 1

			return url
		}, func(ctx context.Context, url *domain.Url) error {
			return nil
		})

	return assertion
}

func arrangeUpdate(t *testing.T) *assert.Assertions {
	ID := uint(10)
	createUrl(&ID)

	factory = New(repository, Options{})
	assertion := assert.New(t)

	repository.
		On(
			"UpdateUrl",
			mock.Anything,
			mock.AnythingOfType("*url.Url"),
		).
		Return(func(ctx context.Context, url *domain.Url) *domain.Url {
			assertion.Equal(mainUrl, url)

			return url
		}, func(ctx context.Context, url *domain.Url) error {
			return nil
		})

	return assertion
}

func arrangeDelete(t *testing.T) *assert.Assertions {
	ID := uint(10)
	createUrl(&ID)

	factory = New(repository, Options{})
	assertion := assert.New(t)

	repository.
		On(
			"DeleteUrl",
			mock.Anything,
			mock.AnythingOfType("uint"),
		).
		Return(func(ctx context.Context, ID uint) error {
			assertion.Equal(mainUrl.ID, ID)

			return nil
		})

	repository.
		On(
			"UrlById",
			mock.Anything,
			mock.AnythingOfType("uint"),
		).
		Return(func(ctx context.Context, ID uint) *domain.Url {
			assertion.Equal(mainUrl.ID, ID)

			return mainUrl
		}, func(ctx context.Context, ID uint) error {
			return nil
		})

	return assertion
}

func arrangeUpsertCreate(t *testing.T) *assert.Assertions {
	createUrl(nil)

	factory = New(repository, Options{})
	assertion := assert.New(t)

	repository.
		On(
			"UpsertUrls",
			mock.Anything,
			mock.Anything,
		).
		Return(func(ctx context.Context, urls ...*domain.Url) []*domain.Url {
			assertion.Equal(mainUrl, urls[0])

			urls[0].ID = 10

			return urls
		}, func(ctx context.Context, urls ...*domain.Url) error {
			return nil
		})

	return assertion
}

func arrangeUpsertUpdate(t *testing.T) *assert.Assertions {
	ID := uint(10)
	createUrl(&ID)

	factory = New(repository, Options{})
	assertion := assert.New(t)

	repository.
		On(
			"UpsertUrls",
			mock.Anything,
			mock.Anything,
		).
		Return(func(ctx context.Context, urls ...*domain.Url) []*domain.Url {
			assertion.Equal(mainUrl, urls[0])

			return urls
		}, func(ctx context.Context, urls ...*domain.Url) error {
			return nil
		})

	return assertion
}

func arrangeList(t *testing.T) *assert.Assertions {
	ID := uint(10)
	createUrl(&ID)

	factory = New(repository, Options{})
	assertion := assert.New(t)

	repository.
		On(
			"ListUrls",
			mock.Anything,
			mock.AnythingOfType("uint"),
			mock.AnythingOfType("queryParameter.QueryParameter"),
		).
		Return(func(ctx context.Context, coinId uint, parameter queryParameter.QueryParameter) []*domain.Url {
			assertion.Equal(mainUrl.CoinID, coinId)

			var urls []*domain.Url
			urls = append(urls, mainUrl)

			return urls
		}, func(ctx context.Context, coinId uint, parameter queryParameter.QueryParameter) error {
			return nil
		})

	return assertion
}

func arrangeCount(t *testing.T) *assert.Assertions {
	ID := uint(10)
	createUrl(&ID)

	factory = New(repository, Options{})
	assertion := assert.New(t)

	repository.
		On(
			"CountUrls",
			mock.Anything,
			mock.AnythingOfType("uint"),
		).
		Return(func(ctx context.Context, coinId uint) uint64 {
			assertion.Equal(mainUrl.CoinID, coinId)

			return 1
		}, func(ctx context.Context, coinId uint) error {
			return nil
		})

	return assertion
}

func createUrl(ID *uint) {
	urlLink, _ := link.New("https://test.com")
	urlType, _ := socialMedia.New("twitter")
	urlExternalId, _ := externalId.New(1)
	mainUrl = domain.New(
		urlExternalId,
		*urlLink,
		*urlType,
	)

	if ID != nil {
		mainUrl.ID = *ID
	}
}

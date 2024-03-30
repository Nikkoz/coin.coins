package url

import (
	coinPrepearer "coins/internal/delivery/http/handlers/coin"
	urlResponse "coins/internal/delivery/http/responses/url"
	domain "coins/internal/domain/url"
	"coins/internal/domain/url/types/link"
	"coins/internal/domain/url/types/socialMedia"
	"coins/pkg/types/context"
	"coins/pkg/types/logger"
	"github.com/gin-gonic/gin"
)

// prepare generates data about the model
func prepare(c *gin.Context, setUrlId bool) (*domain.Url, error) {
	coinId, err := coinPrepearer.Id(c)
	if err != nil {
		return nil, err
	}

	url, err := newUrl(c)
	if err != nil {
		return nil, err
	}

	if setUrlId {
		urlId, err := id(c)
		if err != nil {
			return nil, err
		}

		url.ID = urlId.Value
	}

	url.CoinID = coinId.Value

	return url, nil
}

func id(c *gin.Context) (*urlResponse.ID, error) {
	id := &urlResponse.ID{}
	if err := c.ShouldBindUri(&id); err != nil {
		return nil, logger.ErrorWithContext(context.New(c), err)
	}

	return id, nil
}

func newUrl(c *gin.Context) (*domain.Url, error) {
	url := urlResponse.Short{}
	if err := c.ShouldBindJSON(&url); err != nil {
		return nil, logger.ErrorWithContext(context.New(c), err)
	}

	urlLink, err := link.New(url.Link)
	if err != nil {
		return nil, logger.ErrorWithContext(context.New(c), err)
	}

	urlType, err := socialMedia.New(url.Type)
	if err != nil {
		return nil, logger.ErrorWithContext(context.New(c), err)
	}

	return domain.New(nil, *urlLink, *urlType), nil
}

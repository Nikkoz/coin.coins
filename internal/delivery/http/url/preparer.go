package url

import (
	"coins/internal/delivery/http/coin"
	domain "coins/internal/domain/url"
	"coins/internal/domain/url/types/link"
	"coins/internal/domain/url/types/socialMedia"
	"fmt"
	"github.com/gin-gonic/gin"
)

// prepare generates data about the model
func prepare(c *gin.Context, setUrlId bool) (*domain.Url, error) {
	coinId, err := coin.Id(c)
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

func id(c *gin.Context) (*ID, error) {
	id := &ID{}
	if err := c.ShouldBindUri(&id); err != nil {
		return nil, err
	}

	return id, nil
}

func newUrl(c *gin.Context) (*domain.Url, error) {
	url := Short{}
	if err := c.ShouldBindJSON(&url); err != nil {
		return nil, fmt.Errorf("payload is not correct, Error: %w", err)
	}

	urlLink, err := link.New(url.Link)
	if err != nil {
		return nil, err
	}

	urlType, err := socialMedia.New(url.Type)
	if err != nil {
		return nil, err
	}

	return domain.New(nil, *urlLink, *urlType), nil
}

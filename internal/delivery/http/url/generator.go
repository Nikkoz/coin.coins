package url

import (
	domain "coins/internal/domain/url"
	"coins/internal/domain/url/types/link"
	"coins/internal/domain/url/types/socialMedia"
	"fmt"
	"github.com/gin-gonic/gin"
)

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

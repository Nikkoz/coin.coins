package coin

import (
	domain "coins/internal/domain/coin"
	"coins/internal/domain/coin/types/code"
	"coins/internal/domain/coin/types/name"
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

func newCoin(c *gin.Context) (*domain.Coin, error) {
	coin := Short{}
	if err := c.ShouldBindJSON(&coin); err != nil {
		return nil, fmt.Errorf("payload is not correct, Error: %w", err)
	}

	coinName, err := name.New(coin.Name)
	if err != nil {
		return nil, err
	}

	coinCode, err := code.New(coin.Code)
	if err != nil {
		return nil, err
	}

	return domain.New(*coinName, *coinCode, nil), nil
}

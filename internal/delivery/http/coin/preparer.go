package coin

import (
	domain "coins/internal/domain/coin"
	"coins/internal/domain/coin/types/code"
	"coins/internal/domain/coin/types/name"
	"coins/pkg/types/context"
	"coins/pkg/types/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

func prepare(c *gin.Context, setCoinId bool) (*domain.Coin, error) {
	coin, err := newCoin(c)
	if err != nil {
		return nil, err
	}

	if setCoinId {
		coinId, err := Id(c)
		if err != nil {
			return nil, err
		}

		coin.ID = coinId.Value
	}

	return coin, nil
}

func Id(c *gin.Context) (*ID, error) {
	id := &ID{}
	if err := c.ShouldBindUri(&id); err != nil {
		return nil, logger.ErrorWithContext(context.New(c), err)
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
		return nil, logger.ErrorWithContext(context.New(c), err)
	}

	coinCode, err := code.New(coin.Code)
	if err != nil {
		return nil, logger.ErrorWithContext(context.New(c), err)
	}

	return domain.New(*coinName, *coinCode, nil), nil
}

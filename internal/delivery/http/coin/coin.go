package coin

import (
	"coins/internal/delivery/http/error"
	domain "coins/internal/domain/coin"
	"coins/internal/domain/coin/types/code"
	"coins/internal/domain/coin/types/name"
	useCase "coins/internal/useCase/interfaces"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	useCase useCase.Coin
}

func New(uc useCase.Coin) *Handler {
	return &Handler{
		useCase: uc,
	}
}

func (handler *Handler) Create(c *gin.Context) {
	coin := Short{}
	if err := c.ShouldBindJSON(&coin); err != nil {
		error.SetError(c, http.StatusBadRequest, fmt.Errorf("payload is not correct, Error: %w", err))

		return
	}

	coinName, err := name.New(coin.Name)
	if err != nil {
		error.SetError(c, http.StatusBadRequest, err)

		return
	}

	coinCode, err := code.New(coin.Code)
	if err != nil {
		error.SetError(c, http.StatusBadRequest, err)

		return
	}

	newCoin := domain.New(*coinName, *coinCode, nil)

	response, err := handler.useCase.Create(newCoin)
	if err != nil {
		error.SetError(c, http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusCreated, ToResponse(response))
}

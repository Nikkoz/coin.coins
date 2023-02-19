package coin

import (
	deliveryErr "coins/internal/delivery/http/error"
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
	coin, err := newCoin(c)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return
	}

	response, err := handler.useCase.Create(coin)
	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusCreated, ToResponse(response))
}

func (handler *Handler) Update(c *gin.Context) {
	coinId, err := id(c)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return
	}

	coin, err := newCoin(c)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return
	}

	coin.ID = coinId.Value

	response, err := handler.useCase.Update(coin)
	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, ToResponse(response))
}

func (handler *Handler) Delete(c *gin.Context) {
	coinId, err := id(c)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return
	}

	err = handler.useCase.Delete(coinId.Value)
	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return
	}

	c.Status(http.StatusNoContent)
}

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

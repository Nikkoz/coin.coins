package coin

import (
	deliveryErr "coins/internal/delivery/http/error"
	useCase "coins/internal/useCase/interfaces"
	"coins/pkg/types/pagination"
	"coins/pkg/types/query"
	"coins/pkg/types/queryParameter"
	"github.com/gin-gonic/gin"
	"net/http"
)

var mappingSort = query.SortsOptions{
	"id":   {},
	"name": {},
	"code": {},
}

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

func (handler *Handler) List(c *gin.Context) {
	params, err := query.Parse(c, query.Options{
		Sorts: mappingSort,
	})
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return
	}

	coins, err := handler.useCase.List(queryParameter.QueryParameter{
		Sorts: params.Sorts,
		Pagination: pagination.Pagination{
			Limit:  params.Limit,
			Offset: params.Offset,
		},
	})
	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return
	}

	count, err := handler.useCase.Count()
	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, ToListResponse(count, *params, coins))
}

package coin

import (
	"coins/internal/delivery/http/actions"
	deliveryErr "coins/internal/delivery/http/error"
	domain "coins/internal/domain/coin"
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
	response := handler.save(c, actions.Create)
	if response == nil {
		return
	}

	c.JSON(http.StatusCreated, ToResponse(response))
}

func (handler *Handler) Update(c *gin.Context) {
	response := handler.save(c, actions.Update)
	if response == nil {
		return
	}

	c.JSON(http.StatusOK, ToResponse(response))
}

func (handler *Handler) Delete(c *gin.Context) {
	coinId, err := Id(c)
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
			Limit: params.Limit,
			Page:  params.Page,
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

func (handler *Handler) save(c *gin.Context, action actions.Action) *domain.Coin {
	coin, err := prepare(c, action == actions.Update)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return nil
	}

	var response *domain.Coin

	switch action {
	case actions.Create:
		response, err = handler.useCase.Create(coin)
	case actions.Update:
		response, err = handler.useCase.Update(coin)
	}

	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return nil
	}

	return response
}

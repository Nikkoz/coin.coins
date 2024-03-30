package coin

import (
	deliveryErr "coins/internal/delivery/http/error"
	domain "coins/internal/domain/coin"
	useCase "coins/internal/useCases/interfaces"
	"coins/pkg/types/action"
	"coins/pkg/types/context"
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
	response := handler.save(c, action.Create)
	if response == nil {
		return
	}

	c.JSON(http.StatusCreated, toResponse(response))
}

func (handler *Handler) Update(c *gin.Context) {
	response := handler.save(c, action.Update)
	if response == nil {
		return
	}

	c.JSON(http.StatusOK, toResponse(response))
}

func (handler *Handler) Delete(c *gin.Context) {
	coinId, err := Id(c)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return
	}

	ctx := context.New(c)

	err = handler.useCase.Delete(ctx, coinId.Value)
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

	// @todo: add timeout for context
	ctx := context.New(c)

	coins, err := handler.useCase.List(ctx, queryParameter.QueryParameter{
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

	count, err := handler.useCase.Count(ctx)
	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, toListResponse(count, *params, coins))
}

func (handler *Handler) save(c *gin.Context, act action.Action) *domain.Coin {
	coin, err := prepare(c, act == action.Update)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return nil
	}

	var response *domain.Coin
	var ctx = context.New(c)

	switch act {
	case action.Create:
		response, err = handler.useCase.Create(ctx, coin)
	case action.Update:
		response, err = handler.useCase.Update(ctx, coin)
	}

	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return nil
	}

	return response
}

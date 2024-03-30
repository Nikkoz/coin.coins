package url

import (
	deliveryErr "coins/internal/delivery/http/error"
	coinPreparer "coins/internal/delivery/http/handlers/coin"
	domain "coins/internal/domain/url"
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
	"type": {},
}

type Handler struct {
	useCase useCase.Url
}

func New(uc useCase.Url) *Handler {
	return &Handler{
		useCase: uc,
	}
}

func (handler *Handler) Create(c *gin.Context) {
	response := handler.save(c, action.Create)
	if response == nil {
		return
	}

	c.JSON(http.StatusCreated, ToResponse(response))
}

func (handler *Handler) Update(c *gin.Context) {
	response := handler.save(c, action.Update)
	if response == nil {
		return
	}

	c.JSON(http.StatusOK, ToResponse(response))
}

func (handler *Handler) Delete(c *gin.Context) {
	coinId, err := coinPreparer.Id(c)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return
	}

	urlId, err := id(c)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return
	}

	ctx := context.New(c)

	err = handler.useCase.Delete(ctx, urlId.Value, coinId.Value)
	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return
	}

	c.Status(http.StatusNoContent)
}

func (handler *Handler) List(c *gin.Context) {
	coinId, err := coinPreparer.Id(c)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return
	}

	params, err := query.Parse(c, query.Options{
		Sorts: mappingSort,
	})
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return
	}

	ctx := context.New(c)

	urls, err := handler.useCase.List(ctx, coinId.Value, queryParameter.QueryParameter{
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

	count, err := handler.useCase.Count(ctx, coinId.Value)
	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, ToListResponse(count, *params, urls))
}

func (handler *Handler) save(c *gin.Context, act action.Action) *domain.Url {
	url, err := prepare(c, act == action.Update)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return nil
	}

	var response *domain.Url
	var ctx = context.New(c)

	switch act {
	case action.Create:
		response, err = handler.useCase.Create(ctx, url)
	case action.Update:
		response, err = handler.useCase.Update(ctx, url)
	}

	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return nil
	}

	return response
}

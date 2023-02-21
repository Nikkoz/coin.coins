package url

import (
	"coins/internal/delivery/http/actions"
	"coins/internal/delivery/http/coin"
	deliveryErr "coins/internal/delivery/http/error"
	domain "coins/internal/domain/url"
	useCase "coins/internal/useCase/interfaces"
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
	coinId, err := coin.Id(c)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return
	}

	urlId, err := id(c)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return
	}

	err = handler.useCase.Delete(urlId.Value, coinId.Value)
	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return
	}

	c.Status(http.StatusNoContent)
}

func (handler *Handler) List(c *gin.Context) {
	coinId, err := coin.Id(c)
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

	urls, err := handler.useCase.List(coinId.Value, queryParameter.QueryParameter{
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

	count, err := handler.useCase.Count(coinId.Value)
	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, ToListResponse(count, *params, urls))
}

func (handler *Handler) save(c *gin.Context, action actions.Action) *domain.Url {
	url, err := prepare(c, action == actions.Update)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return nil
	}

	var response *domain.Url

	switch action {
	case actions.Create:
		response, err = handler.useCase.Create(url)
	case actions.Update:
		response, err = handler.useCase.Update(url)
	}

	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return nil
	}

	return response
}

package url

import (
	"coins/internal/delivery/http/actions"
	deliveryErr "coins/internal/delivery/http/error"
	domain "coins/internal/domain/url"
	useCase "coins/internal/useCase/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

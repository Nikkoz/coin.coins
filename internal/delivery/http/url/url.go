package url

import (
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

type action string

const (
	Create action = "create"
	Update action = "update"
)

func (handler *Handler) Create(c *gin.Context) {
	response := handler.save(c, Create)
	if response == nil {
		return
	}

	c.JSON(http.StatusCreated, ToResponse(response))
}

func (handler *Handler) Update(c *gin.Context) {
	response := handler.save(c, Update)
	if response == nil {
		return
	}

	c.JSON(http.StatusOK, ToResponse(response))
}

func (handler *Handler) save(c *gin.Context, action action) *domain.Url {
	u, err := prepare(c, action == Update)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return nil
	}

	var response *domain.Url

	switch action {
	case Create:
		response, err = handler.useCase.Create(u)
	case Update:
		response, err = handler.useCase.Update(u)
	}

	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return nil
	}

	return response
}

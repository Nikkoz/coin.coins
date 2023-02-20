package url

import (
	"coins/internal/delivery/http/coin"
	deliveryErr "coins/internal/delivery/http/error"
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
	ID, err := coin.Id(c)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return
	}

	url, err := newUrl(c)
	if err != nil {
		deliveryErr.SetError(c, http.StatusBadRequest, err)

		return
	}

	url.CoinID = ID.Value

	response, err := handler.useCase.Create(url)
	if err != nil {
		deliveryErr.SetError(c, http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusCreated, ToResponse(response))
}

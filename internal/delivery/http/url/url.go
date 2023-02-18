package url

import (
	useCase "coins/internal/useCase/interfaces"
	"github.com/gin-gonic/gin"
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

}

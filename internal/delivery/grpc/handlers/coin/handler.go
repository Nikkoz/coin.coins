package coin

import (
	"coins/internal/domain/coin"
	useCase "coins/internal/useCase/interfaces"
	"coins/pkg/protobuf/coins"
	"coins/pkg/types/context"
	"coins/pkg/types/pagination"
	"coins/pkg/types/queryParameter"
)

type Handler struct {
	useCase useCase.Coin
	topic   string
}

func New(uc useCase.Coin, topic string) *Handler {
	return &Handler{
		useCase: uc,
		topic:   topic,
	}
}

func (h *Handler) List(ctx context.Context, page uint64) ([]*coin.Coin, error) {
	models, err := h.useCase.List(ctx, queryParameter.QueryParameter{
		Pagination: pagination.Pagination{
			Page: page,
		},
	})
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (h *Handler) ByID(ctx context.Context, id uint64) (*coins.Coin, error) {
	model, c, err := h.useCase.FindFullByID(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	return modelToPb(model, c, true), nil
}

package grpc

import (
	"coins/internal/domain/coin"
	"coins/pkg/protobuf/coins"
)

func modelsToPb(models []*coin.Coin) []*coins.Coin {
	pbCoins := make([]*coins.Coin, len(models))
	for k, model := range models {
		pbCoins[k] = modelToPb(model, false)
	}

	return pbCoins
}

func modelToPb(model *coin.Coin, full bool) *coins.Coin {
	c := &coins.Coin{
		Id:   uint64(model.ID),
		Name: model.Name.String(),
		Code: model.Code.String(),
		Icon: model.Icon.String(),
	}

	if full {
		c.Info = &coins.Info{}
	}

	return c
}

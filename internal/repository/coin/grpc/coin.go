package grpc

import (
	"coins/pkg/types/context"
	"coins/pkg/types/logger"
	"github.com/Nikkoz/coin.sync/pkg/protobuf/coins"
)

func (r *Repository) GetInfo(ctx context.Context, id uint64) (*coins.Coin, error) {
	response, err := r.GetInfos(ctx, []uint64{id})
	if err != nil {
		return nil, err
	}

	return response[0], nil
}

func (r *Repository) GetInfos(ctx context.Context, ids []uint64) ([]*coins.Coin, error) {
	response, err := r.client.GetCoins(ctx, ToCoinRequest(ids))
	if err != nil {
		return nil, logger.ErrorWithContext(ctx, err)
	}

	return response.GetCoins(), nil
}

package grpc

import (
	"coins/pkg/protobuf/coins"
	cpk "coins/pkg/types/context"
	"context"
	"time"
)

func (d *Delivery) GetCoins(c context.Context, request *coins.GetCoinsRequest) (*coins.GetCoinsResponse, error) {
	ctx := cpk.New(c)

	models, err := d.Handlers.Coin.List(ctx, request.GetPage())
	if err != nil {
		return nil, err
	}

	return &coins.GetCoinsResponse{Coins: modelsToPb(models)}, nil
}

func (d *Delivery) GetCoin(c context.Context, request *coins.GetCoinRequest) (*coins.GetCoinResponse, error) {
	ctx := cpk.NewWithTimeout(c, time.Second*60)

	model, err := d.Handlers.Coin.ByID(ctx, request.GetId())
	if err != nil {
		return nil, err
	}

	return &coins.GetCoinResponse{Coin: model}, nil
}

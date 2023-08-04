package grpc

import (
	"coins/pkg/types/context"
	"github.com/Nikkoz/coin.sync/pkg/protobuf/coins"
)

type Coin interface {
	GetInfo(ctx context.Context, id uint64) (*coins.Coin, error)
	GetInfos(ctx context.Context, ids []uint64) ([]*coins.Coin, error)
}

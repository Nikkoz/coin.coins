package grpc

import (
	"coins/internal/repositories/coin/interfaces"
	"github.com/Nikkoz/coin.sync/pkg/protobuf/coins"
	"google.golang.org/grpc"
)

var _ interfaces.Grpc = (*Repository)(nil)

type (
	Repository struct {
		client  coins.CoinServiceClient
		options Options
	}

	Options struct{}
)

func New(conn *grpc.ClientConn, o Options) *Repository {
	repo := &Repository{
		client: coins.NewCoinServiceClient(conn),
	}

	repo.SetOptions(o)

	return repo
}

func (r *Repository) SetOptions(options Options) {
	if r.options != options {
		r.options = options
	}
}

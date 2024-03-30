package grpc

import (
	"coins/configs"
	grpcClient "coins/pkg/grpc"
	"coins/pkg/types/context"
	"coins/pkg/types/logger"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type ServiceProvider struct {
}

func New(ctx context.Context, config configs.Config) *grpc.ClientConn {
	connection, err := grpcClient.New(ctx, config.Grpc.Host, config.Grpc.Port, config.App.Name, config.App.Version).GetConnection()
	if err != nil {
		_ = logger.ErrorWithContext(ctx, errors.Wrap(err, "failed to create grpc client"))
	}

	return connection
}

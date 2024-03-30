package coin

import (
	repo "coins/internal/repositories/coin/interfaces"
	"coins/internal/useCases/interfaces"
)

var _ interfaces.Coin = (*Factory)(nil)

type (
	Factory struct {
		storage repo.Storage
		broker  repo.Broker
		grpc    repo.Grpc

		options Options
	}

	Options struct{}
)

func New(s repo.Storage, b repo.Broker, g repo.Grpc, options Options) *Factory {
	factory := &Factory{
		storage: s,
		broker:  b,
		grpc:    g,
	}

	factory.SetOption(options)

	return factory
}

func (factory *Factory) SetOption(options Options) {
	if factory.options != options {
		factory.options = options
	}
}

package coin

import (
	"coins/internal/useCase/adapters/broker"
	"coins/internal/useCase/adapters/grpc"
	"coins/internal/useCase/adapters/storage"
)

type (
	Factory struct {
		adapterStorage storage.Coin
		adapterBroker  broker.Coin
		adapterGrpc    grpc.Coin

		options Options
	}

	Options struct{}
)

func New(s storage.Coin, b broker.Coin, g grpc.Coin, options Options) *Factory {
	factory := &Factory{
		adapterStorage: s,
		adapterBroker:  b,
		adapterGrpc:    g,
	}

	factory.SetOption(options)

	return factory
}

func (factory *Factory) SetOption(options Options) {
	if factory.options != options {
		factory.options = options
	}
}

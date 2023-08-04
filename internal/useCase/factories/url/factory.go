package url

import "coins/internal/useCase/adapters/storage"

type (
	Factory struct {
		adapterStorage storage.Url
		options        Options
	}

	Options struct{}
)

func New(s storage.Url, options Options) *Factory {
	factory := &Factory{
		adapterStorage: s,
	}

	factory.SetOption(options)

	return factory
}

func (f *Factory) SetOption(options Options) {
	if f.options != options {
		f.options = options
	}
}

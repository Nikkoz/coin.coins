package url

import (
	repo "coins/internal/repositories/url/interfaces"
	"coins/internal/useCases/interfaces"
)

var _ interfaces.Url = (*Factory)(nil)

type (
	Factory struct {
		storage repo.Storage
		options Options
	}

	Options struct{}
)

func New(s repo.Storage, options Options) *Factory {
	factory := &Factory{
		storage: s,
	}

	factory.SetOption(options)

	return factory
}

func (f *Factory) SetOption(options Options) {
	if f.options != options {
		f.options = options
	}
}

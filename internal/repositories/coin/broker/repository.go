package broker

import (
	"coins/internal/repositories/coin/interfaces"
	"coins/pkg/store/messageBroker"
)

var _ interfaces.Broker = (*Repository)(nil)

type (
	Repository struct {
		broker  messageBroker.MessageBroker
		options Options
	}

	Options struct{}
)

func New(broker messageBroker.MessageBroker, options Options) *Repository {
	repo := &Repository{
		broker: broker,
	}

	repo.SetOptions(options)

	return repo
}

func (r *Repository) SetOptions(options Options) {
	if r.options != options {
		r.options = options
	}
}

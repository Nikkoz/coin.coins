package broker

import (
	"coins/pkg/store/messageBroker"
)

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

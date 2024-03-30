package broker

import (
	useCase "coins/internal/useCases/interfaces"
	"coins/pkg/store/messageBroker"
	"coins/pkg/types/context"
	"coins/pkg/types/logger"
)

type (
	Delivery struct {
		ucCoin useCase.Coin
		ucUrl  useCase.Url

		options Options
	}

	Options struct {
		Notify chan error
	}
)

func New(ucCoin useCase.Coin, ucUrl useCase.Url, o Options) *Delivery {
	d := &Delivery{
		ucCoin: ucCoin,
		ucUrl:  ucUrl,
	}

	d.setOptions(o)

	return d
}

func (d *Delivery) setOptions(options Options) {
	if options.Notify == nil {
		d.options.Notify = make(chan error, 1)
	}

	if d.options != options {
		d.options = options
	}
}

func (d *Delivery) Run(broker messageBroker.MessageBroker, topics []string) {
	ctx := context.New(context.Empty())

	go func() {
		defer close(d.options.Notify)

		d.options.Notify <- d.ucCoin.Subscribe(ctx, topics)
	}()

	logger.Info("message broker started successfully")
	go func() {
		defer close(d.options.Notify)

		broker.Consume(d.options.Notify, ctx, d.ucCoin.Consume)
	}()
}

func (d *Delivery) Notify() <-chan error {
	return d.options.Notify
}

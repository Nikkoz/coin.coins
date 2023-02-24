package broker

import (
	useCase "coins/internal/useCase/interfaces"
	"coins/pkg/store/messageBroker"
	"coins/pkg/types/context"
	"coins/pkg/types/logger"
	"fmt"
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

	err := d.ucCoin.Subscribe(ctx, topics)
	if err != nil {
		d.options.Notify <- fmt.Errorf("can't subscribe on topics: %v\n", err)
		close(d.options.Notify)

		return
	}

	logger.Info("message broker started successfully")
	go broker.Consume(d.options.Notify, ctx, d.ucCoin.Consume)
}

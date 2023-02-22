package broker

import (
	useCase "coins/internal/useCase/interfaces"
	"coins/pkg/store/messageBroker"
	"coins/pkg/types/context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type (
	Delivery struct {
		ucCoin useCase.Coin
		ucUrl  useCase.Url

		options Options
	}

	Options struct{}
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
	if d.options != options {
		d.options = options
	}
}

func (d *Delivery) Run(broker messageBroker.MessageBroker, topics []string) error {
	sigChan := make(chan os.Signal, 1)
	doneChan := make(chan bool)
	ctx := context.New(context.Empty())

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	err := d.ucCoin.Subscribe(ctx, topics)
	if err != nil {
		return fmt.Errorf("can't subscribe on topics: %v\n", err)
	}

	fmt.Println("message broker started successfully")
	go broker.Consume(sigChan, doneChan, ctx, d.ucCoin.Consume)

	<-doneChan

	return nil
}

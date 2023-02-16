package broker

import (
	useCase "coins/internal/useCase/interfaces"
	"coins/pkg/store/messageBroker"
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

func (d *Delivery) Run(broker messageBroker.MessageBroker) error {
	sigChan := make(chan os.Signal, 1)
	doneChan := make(chan bool)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// @todo: продумать как указывать топики
	err := d.ucCoin.Subscribe([]string{"upload_coins"})
	if err != nil {
		return fmt.Errorf("can't subscribe on topics: %v\n", err)
	}

	go broker.Consume(sigChan, doneChan, d.ucCoin.Consume)

	<-doneChan

	return nil
}

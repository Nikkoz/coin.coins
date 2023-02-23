package http

import (
	"coins/configs"
	coinHandler "coins/internal/delivery/http/coin"
	urlHandler "coins/internal/delivery/http/url"
	useCase "coins/internal/useCase/interfaces"
	"fmt"
	"github.com/gin-gonic/gin"
)

type (
	Delivery struct {
		ucCoin useCase.Coin
		ucUrl  useCase.Url

		router *gin.Engine

		options  Options
		Handlers Handlers
	}

	Options struct {
		Notify chan error
	}

	Handlers struct {
		Coin *coinHandler.Handler
		Url  *urlHandler.Handler
	}
)

func New(ucCoin useCase.Coin, ucUrl useCase.Url, o Options) *Delivery {
	d := &Delivery{
		ucCoin: ucCoin,
		ucUrl:  ucUrl,
	}

	d.setOptions(o)
	d.setHandlers()

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

func (d *Delivery) setHandlers() {
	d.Handlers = Handlers{
		Coin: coinHandler.New(d.ucCoin),
		Url:  urlHandler.New(d.ucUrl),
	}
}

func (d *Delivery) Run(config configs.Config) {
	d.initRouter(config)

	go func() {
		if err := d.router.Run(fmt.Sprintf("%s:%d", config.Http.Host, config.Http.Port)); err != nil {
			d.options.Notify <- err

			close(d.options.Notify)
		}
	}()
}

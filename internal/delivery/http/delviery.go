package http

import (
	"coins/configs"
	coinHandler "coins/internal/delivery/http/coin"
	urlHandler "coins/internal/delivery/http/url"
	useCase "coins/internal/useCase/interfaces"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

type (
	Delivery struct {
		ucCoin useCase.Coin
		ucUrl  useCase.Url

		router *gin.Engine

		options  Options
		Handlers Handlers
	}

	Options struct{}

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

func (d *Delivery) Run(config configs.Config) (err error) {
	d.initRouter(config)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err = d.router.Run(fmt.Sprintf("%s:%d", config.Http.Host, config.Http.Port))
	}()

	<-signalCh

	return err
}

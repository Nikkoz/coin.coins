package grpc

import (
	"coins/configs"
	coinHandler "coins/internal/delivery/grpc/handlers/coin"
	"coins/internal/useCases/interfaces"
	"coins/pkg/protobuf/coins"
	"coins/pkg/types/logger"
	"flag"
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcCtxTags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpcOpentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

type (
	Delivery struct {
		coins.UnimplementedCoinServiceServer

		options  Options
		Handlers Handlers
	}

	Options struct {
		Notify      chan error
		BrokerTopic string
	}

	Handlers struct {
		Coin *coinHandler.Handler
	}
)

func New(ucCoin interfaces.Coin, o Options) *Delivery {
	d := &Delivery{}

	d.setOptions(o)
	d.setHandlers(ucCoin)

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

func (d *Delivery) setHandlers(ucCoin interfaces.Coin) {
	d.Handlers = Handlers{
		Coin: coinHandler.New(ucCoin, d.options.BrokerTopic),
	}
}

func (d *Delivery) Run(config configs.Config) (*grpc.Server, net.Listener) {
	flag.Parse()

	grpcAddr := fmt.Sprintf("%s:%v", config.Grpc.Host, config.Grpc.Port)

	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		logger.Error(fmt.Errorf("failed to listen: %w", err))

		return nil, nil
	}

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Duration(config.Grpc.MaxConnectionIdle) * time.Second,
			Timeout:           time.Duration(config.Grpc.Timeout) * time.Second,
			MaxConnectionAge:  time.Duration(config.Grpc.MaxConnectionAge) * time.Second,
			Time:              time.Duration(config.Grpc.Timeout) * time.Minute,
		}),
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcCtxTags.UnaryServerInterceptor(),
			grpcOpentracing.UnaryServerInterceptor(),
			grpcRecovery.UnaryServerInterceptor(),
		)),
	)

	coins.RegisterCoinServiceServer(grpcServer, d)

	go func() {
		logger.Info(fmt.Sprintf("GRPC Server is listening on: %s", grpcAddr))

		if err := grpcServer.Serve(listener); err != nil {
			d.options.Notify <- fmt.Errorf("failed running gRPC server: %w", err)
		}
	}()

	return grpcServer, listener
}

func (d *Delivery) Notify() <-chan error {
	return d.options.Notify
}

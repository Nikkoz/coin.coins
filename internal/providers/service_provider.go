package providers

import (
	"coins/configs"
	brokerProvider "coins/internal/providers/broker"
	"coins/internal/providers/delivery"
	grpcProvider "coins/internal/providers/grpc"
	"coins/internal/providers/repository"
	coinFactory "coins/internal/useCases/factories/coin"
	urlFactory "coins/internal/useCases/factories/url"
	"coins/pkg/store/messageBroker"
	"coins/pkg/types/context"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ServiceProvider struct {
	ctx    context.Context
	config configs.Config
	db     *gorm.DB

	Broker   messageBroker.MessageBroker
	Grpc     *grpc.ClientConn
	Delivery *delivery.ServiceProvider
}

func New(ctx context.Context, config configs.Config, db *gorm.DB) *ServiceProvider {
	provider := &ServiceProvider{
		ctx:    ctx,
		config: config,
		db:     db,
	}

	provider.init()

	return provider
}

func (s *ServiceProvider) init() {
	s.Broker = brokerProvider.New(s.config)
	s.Grpc = grpcProvider.New(s.ctx, s.config)

	s.deliveryInit()
}

func (s *ServiceProvider) deliveryInit() {
	repositoryProvider := repository.New(s.db)
	cf := coinFactory.New(
		repositoryProvider.Coin(),
		repositoryProvider.Broker(s.Broker),
		repositoryProvider.Grpc(s.Grpc),
		coinFactory.Options{},
	)
	uf := urlFactory.New(repositoryProvider.Url(), urlFactory.Options{})

	s.Delivery = delivery.New(s.db, cf, uf)
}

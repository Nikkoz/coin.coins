package delivery

import (
	"coins/internal/delivery/broker"
	"coins/internal/delivery/grpc"
	"coins/internal/delivery/http"
	"coins/internal/useCases/factories/coin"
	"coins/internal/useCases/factories/url"
	"gorm.io/gorm"
)

type ServiceProvider struct {
	connection *gorm.DB

	Messenger *broker.Delivery
	Http      *http.Delivery
	Grpc      *grpc.Delivery
}

func New(connection *gorm.DB, cf *coin.Factory, uf *url.Factory) *ServiceProvider {
	provider := &ServiceProvider{
		connection: connection,
	}

	provider.init(cf, uf)

	return provider
}

func (s *ServiceProvider) init(cf *coin.Factory, uf *url.Factory) {
	s.messengerInit(cf, uf)
	s.httpInit(cf, uf)
	s.grpcInit(cf)
}

func (s *ServiceProvider) messengerInit(cf *coin.Factory, uf *url.Factory) {
	if s.Messenger == nil {
		s.Messenger = broker.New(cf, uf, broker.Options{})
	}
}

func (s *ServiceProvider) httpInit(cf *coin.Factory, uf *url.Factory) {
	if s.Http == nil {
		s.Http = http.New(cf, uf, http.Options{})
	}
}

func (s *ServiceProvider) grpcInit(cf *coin.Factory) {
	if s.Grpc == nil {
		s.Grpc = grpc.New(cf, grpc.Options{})
	}
}

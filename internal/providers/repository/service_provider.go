package repository

import (
	"coins/internal/repositories/coin/broker"
	coin "coins/internal/repositories/coin/database"
	"coins/internal/repositories/coin/grpc"
	url "coins/internal/repositories/url/database"
	"coins/pkg/store/messageBroker"
	grpcClient "google.golang.org/grpc"
	"gorm.io/gorm"
)

type ServiceProvider struct {
	connection *gorm.DB

	url    *url.Repository
	coin   *coin.Repository
	broker *broker.Repository
	grpc   *grpc.Repository
}

func New(connection *gorm.DB) *ServiceProvider {
	return &ServiceProvider{
		connection: connection,
	}
}

func (s *ServiceProvider) Url() *url.Repository {
	if s.url == nil {
		s.url = url.New(s.connection, url.Options{})
	}

	return s.url
}

func (s *ServiceProvider) Coin() *coin.Repository {
	if s.coin == nil {
		s.coin = coin.New(s.connection, s.Url(), coin.Options{})
	}

	return s.coin
}

func (s *ServiceProvider) Broker(messageBroker messageBroker.MessageBroker) *broker.Repository {
	if s.broker == nil {
		s.broker = broker.New(messageBroker, broker.Options{})
	}

	return s.broker
}

func (s *ServiceProvider) Grpc(conn *grpcClient.ClientConn) *grpc.Repository {
	if s.grpc == nil {
		s.grpc = grpc.New(conn, grpc.Options{})
	}

	return s.grpc
}

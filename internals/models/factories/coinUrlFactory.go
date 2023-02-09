package factories

import (
	"coins/internals/models"
	"coins/internals/models/repositories"
	"gorm.io/gorm"
)

type CoinUrlFactory struct {
	conn       *gorm.DB
	repository *repositories.CoinUrlRepository
}

func NewCoinUrlFactory(conn *gorm.DB) *CoinUrlFactory {
	return &CoinUrlFactory{
		conn:       conn,
		repository: repositories.NewCoinUrlRepository(conn),
	}
}

func (factory *CoinUrlFactory) Upsert(urls []models.CoinUrl) error {
	return factory.repository.Upsert(urls).Error
}

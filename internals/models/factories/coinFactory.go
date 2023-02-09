package factories

import (
	"coins/internals/models"
	"coins/internals/models/repositories"
	"gorm.io/gorm"
)

type CoinFactory struct {
	conn       *gorm.DB
	repository *repositories.CoinRepository
}

func NewCoinFactory(conn *gorm.DB) *CoinFactory {
	return &CoinFactory{
		conn:       conn,
		repository: repositories.NewCoinRepository(conn),
	}
}

func (factory *CoinFactory) Upsert(coins []models.Coin) (err error) {
	err = factory.conn.Transaction(func(tx *gorm.DB) error {
		if err := factory.repository.UpsertWithoutAssociations(coins).Error; err != nil {
			return err
		}

		return factory.saveAssociations(coins)
	})

	return
}

func (factory *CoinFactory) saveAssociations(coins []models.Coin) error {
	urls := make([]models.CoinUrl, 0)

	for _, coin := range coins {
		if len(coin.CoinUrls) == 0 {
			return nil
		}

		for _, link := range coin.CoinUrls {
			link.CoinID = coin.ID

			urls = append(urls, *link)
		}
	}

	if err := NewCoinUrlFactory(factory.conn).Upsert(urls); err != nil {
		return err
	}

	return nil
}

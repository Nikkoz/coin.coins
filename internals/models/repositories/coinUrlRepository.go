package repositories

import (
	"coins/internals/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CoinUrlRepository struct {
	conn *gorm.DB
}

func NewCoinUrlRepository(conn *gorm.DB) *CoinUrlRepository {
	return &CoinUrlRepository{
		conn: conn,
	}
}

func (repository *CoinUrlRepository) Upsert(urls []models.CoinUrl) *gorm.DB {
	return repository.
		conn.
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "external_id"}},
			UpdateAll: true,
		}).
		Create(urls)
}

package repositories

import (
	"coins/internals/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CoinRepository struct {
	conn *gorm.DB
}

func NewCoinRepository(conn *gorm.DB) *CoinRepository {
	return &CoinRepository{
		conn: conn,
	}
}

func (repository *CoinRepository) UpsertWithoutAssociations(coins []models.Coin) *gorm.DB {
	return repository.
		conn.
		Omit(clause.Associations).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).
		Create(coins)
}

package app

import (
	"coins/internal/domain/coin"
	"coins/internal/domain/url"
	"gorm.io/gorm"
	"log"
)

func Migrate(connection *gorm.DB) {
	err := connection.AutoMigrate(
		&coin.Coin{},
		&url.Url{},
	)
	if err != nil {
		log.Fatalf("Error init models: %v\n", err)
	}
}

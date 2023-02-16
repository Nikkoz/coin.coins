package app

import (
	"coins/configs"
	"coins/internal/domain/coin"
	"coins/internal/domain/url"
	"coins/pkg/store/db"
	"fmt"
	"gorm.io/gorm"
	"log"
)

func connectionDB() (*gorm.DB, func()) {
	conn, err := db.New(settingsDB(config.Db))
	if err != nil {
		log.Fatalf("db error: %v", err)
	}

	return conn, func() {
		sqlDB, _ := conn.DB()
		err := sqlDB.Close()
		if err != nil {
			log.Println(fmt.Errorf("Error close connection: %v\n", err))
		}
	}
}

func settingsDB(config configs.Db) db.Settings {
	sslMode := "disable"
	if config.SslMode {
		sslMode = "enable"
	}

	return db.NewSettings(
		config.Connection,
		config.Host,
		config.Port,
		config.Name,
		config.User,
		config.Password,
		sslMode,
		"c_",
		4,
		1000,
	)
}

func Migrate(connection *gorm.DB) {
	err := connection.AutoMigrate(
		&coin.Coin{},
		&url.Url{},
	)
	if err != nil {
		log.Fatalf("Error init models: %v\n", err)
	}
}

package app

import (
	"coins/configs"
	"coins/pkg/store"
	"fmt"
	"gorm.io/gorm"
	"log"
)

var config *configs.Config

func init() {
	cfg, err := configs.New()
	if err != nil {
		log.Fatalf("unable to parse ennvironment variables: %e", err)
	}

	config = cfg
}

func Run() {
	// @todo: see https://github.com/evrone/go-clean-template/blob/34844d644b3cd20696b7bebbec32b0a65678ba7a/internal/app/app.go
	//log := logger.New(config.Log.Level)

	conn, conClose := connection()
	defer conClose()

	Migrate(conn)
}

func connection() (*gorm.DB, func()) {
	conn, err := store.NewConn(settings(config.Db))
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

func settings(config configs.Db) store.Settings {
	sslMode := "disable"
	if config.SslMode {
		sslMode = "enable"
	}

	return store.NewSettings(
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

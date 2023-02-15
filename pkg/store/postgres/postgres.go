package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func NewConn(settings Settings) (*gorm.DB, error) {
	connection, err := gorm.Open(postgres.Open(settings.toDNS()), settings.config())
	if err != nil {
		_, err = fmt.Printf("Can't open connection: %v", err)

		return nil, err
	}

	sqlDB, err := connection.DB()
	if err != nil {
		_, err = fmt.Printf("Can't get sql db: %v", err)

		return nil, err
	}

	sqlDB.SetMaxIdleConns(settings.MaxConnections)
	sqlDB.SetMaxOpenConns(settings.MaxConnections)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return connection, nil
}

package db

import (
	"fmt"
	driver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySql(settings Settings) (*gorm.DB, error) {
	connection, err := gorm.Open(driver.Open(toDNSForMySql(settings)), settings.Config())
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
	sqlDB.SetConnMaxLifetime(settings.LifeTime)

	return connection, nil
}

func toDNSForMySql(settings Settings) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		settings.User,
		settings.Password,
		settings.Host,
		settings.Port,
		settings.Database,
	)
}

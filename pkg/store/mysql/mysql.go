package mysql

import (
	"coins/pkg/store"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func NewConn(settings store.Settings) (*gorm.DB, error) {
	connection, err := gorm.Open(mysql.Open(toDNS(settings)), settings.Config())
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

func toDNS(settings store.Settings) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		settings.User,
		settings.Password,
		settings.Host,
		settings.Port,
		settings.Database,
	)
}

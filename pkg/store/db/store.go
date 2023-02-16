package db

import (
	"coins/pkg/store/db/mysql"
	"coins/pkg/store/db/postgres"
	"errors"
	"gorm.io/gorm"
)

const (
	PGSQL = "pgsql"
	MYSQL = "mysql"
)

var ErrConnectNotSupport = errors.New("type connection not supported")

func New(settings Settings) (*gorm.DB, error) {
	switch settings.Connection {
	case PGSQL:
		return postgres.New(settings)
	case MYSQL:
		return mysql.New(settings)
	default:
		return nil, ErrConnectNotSupport
	}
}

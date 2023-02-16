package db

import (
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
		return NewPgSql(settings)
	case MYSQL:
		return NewMySql(settings)
	default:
		return nil, ErrConnectNotSupport
	}
}

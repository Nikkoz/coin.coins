package store

import (
	"coins/pkg/store/mysql"
	"coins/pkg/store/postgres"
	"errors"
	"gorm.io/gorm"
)

const PGSQL = "pgsql"
const MYSQL = "mysql"

func NewConn(settings Settings) (*gorm.DB, error) {
	switch settings.Connection {
	case PGSQL:
		return postgres.NewConn(settings)
	case MYSQL:
		return mysql.NewConn(settings)
	default:
		return nil, errors.New("type connection not supported")
	}
}

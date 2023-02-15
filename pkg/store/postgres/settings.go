package postgres

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Settings struct {
	Host           string
	Port           uint16
	Database       string
	User           string
	Password       string
	SSLMode        string
	MaxConnections int
	Prefix         string
	BatchSize      int
}

func NewSettings(host string, port uint16, db, user, pass, sslMode, prefix string, maxConn, batchSize int) Settings {
	return Settings{
		Host:           host,
		Port:           port,
		Database:       db,
		User:           user,
		Password:       pass,
		SSLMode:        sslMode,
		Prefix:         prefix,
		MaxConnections: maxConn,
		BatchSize:      batchSize,
	}
}

func (s Settings) toDNS() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		s.Host,
		s.Port,
		s.User,
		s.Password,
		s.Database,
		s.SSLMode,
	)
}

func (s Settings) config() *gorm.Config {
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   s.Prefix,
			SingularTable: false,
		},
		CreateBatchSize: s.BatchSize,
		//DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction: true,
		//DryRun: true,
	}
}

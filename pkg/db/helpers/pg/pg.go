package pg

import (
	"coins/pkg/db"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func Open(cfg db.Config) *gorm.DB {
	dns := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DbName,
		"disable",
	)

	connection, err := gorm.Open(postgres.Open(dns), getGormConfig())
	if err != nil {
		log.Fatalf("Can't open connection: %v", err)
	}

	sqlDB, err := connection.DB()
	if err != nil {
		log.Fatalf("Can't get sql db: %v", err)
	}

	sqlDB.SetMaxOpenConns(4)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return connection
}

func getGormConfig() *gorm.Config {
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "c_",
			SingularTable: false,
		},
		CreateBatchSize: 1000,
		//DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction: true,
		//DryRun: true,
	}
}

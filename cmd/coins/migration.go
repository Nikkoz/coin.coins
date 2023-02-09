package main

import (
	"coins/internals/models"
	"coins/pkg/db"
	"coins/pkg/db/helpers/pg"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		err = fmt.Errorf("Load environment failed: %v\n", err)

		log.Fatal(err)
	}
}

func main() {
	config := db.NewConfig()
	connection := pg.Open(config)

	sqlDB, _ := connection.DB()
	defer func(sqlDb *sql.DB) {
		err := sqlDB.Close()
		if err != nil {
			log.Println(fmt.Errorf("Error close connection: %v\n", err))
		}
	}(sqlDB)

	err := models.InitModels(connection)
	if err != nil {
		log.Println(fmt.Errorf("Error init models: %v\n", err))

		return
	}
}

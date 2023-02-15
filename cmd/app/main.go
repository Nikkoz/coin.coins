package main

import (
	"coins/internal/app"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("load environment failed: %v\n", err)
	}
}

func main() {
	app.Run()
}

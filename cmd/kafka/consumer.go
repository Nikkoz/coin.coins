package main

import (
	"coins/internals/handlers"
	"coins/pkg/kafka"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		err = fmt.Errorf("Load environment failed: %v\n", err)

		log.Fatal(err)
	}
}

func main() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	k, err := kafka.NewKafka()
	defer k.Close()

	if err != nil {
		log.Fatalf("create kafka consumer error: %v", err)

		return
	}

	err = k.Subscribe([]string{"upload_coins"})
	if err != nil {
		log.Fatalf("can't subscribe on topics: %v", err)

		return
	}

	k.Consume(sigchan, handlers.Consume)
}

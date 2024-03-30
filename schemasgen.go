package Coin

//go:generate mkdir -p ./internal/repositories/coin/broker/entities
//go:generate $GOPATH/bin/gogen-avro -package entities -containers ./internal/repositories/coin/broker/entities schemas/coins.avsc

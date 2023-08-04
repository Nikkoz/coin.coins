package Coin

//go:generate mkdir -p ./internal/repository/coin/broker/entities
//go:generate $GOPATH/bin/gogen-avro -package entities -containers ./internal/repository/coin/broker/entities schemas/coins.avsc

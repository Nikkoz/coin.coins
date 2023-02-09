package kafka

//go:generate mkdir -p ../../../internals/entities
//go:generate $GOPATH/bin/gogen-avro -package entities -containers ../../../internals/entities coins.avsc

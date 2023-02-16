package messageBroker

import (
	"errors"
)

const (
	KAFKA    = "kafka"
	RABBITMQ = "rabbitmq"
)

var ErrConnectNotSupport = errors.New("type connection not supported")

func New(settings Settings) (MessageBroker, error) {
	switch settings.Connection {
	case KAFKA:
		return NewKafka(settings)
	case RABBITMQ:
		return nil, ErrConnectNotSupport
	default:
		return nil, ErrConnectNotSupport
	}
}

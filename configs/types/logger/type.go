package logger

type LogLevel string

const (
	Debug LogLevel = "debug"
	Info  LogLevel = "info"
)

func (l LogLevel) IsDebug() bool {
	return l == Debug
}

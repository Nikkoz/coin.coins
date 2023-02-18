package configs

import (
	"coins/configs/types/environment"
	"coins/configs/types/logger"
	"github.com/caarlos0/env/v7"
)

type (
	Config struct {
		App            App            `envPrefix:"APP_"`
		Http           Http           `envPrefix:"HTTP_"`
		Db             Db             `envPrefix:"DB_"`
		Broker         Broker         `envPrefix:"MB_"`
		Sasl           Sasl           `envPrefix:"SASL_"`
		SchemaRegistry SchemaRegistry `envPrefix:"SR_"`
		Log            Log            `envPrefix:"LOG_"`
	}

	App struct {
		Name        string                  `env:"NAME,required"`
		Version     string                  `env:"VERSION,required"`
		Environment environment.Environment `env:"ENV" envDefault:"local"`
	}

	Http struct {
		Host string `env:"HOST" envDefault:"localhost"`
		Port uint16 `env:"PORT" envDefault:"8080"`
	}

	Db struct {
		Connection string `env:"CONNECTION" envDefault:"pgsql"`
		Host       string `env:"HOST" envDefault:"localhost"`
		Port       uint16 `env:"PORT" envDefault:"5432"`
		Name       string `env:"NAME,required"`
		User       string `env:"USER,required"`
		Password   string `env:"PASSWORD,required"`
		SslMode    bool   `env:"USE_SSL" envDefault:"false"`
	}

	Broker struct {
		Connection     string   `env:"CONNECTION" envDefault:"kafka"`
		Broker         string   `env:"BROKER,required"`
		SessionTimeout uint16   `env:"SESSION_TIMEOUT" envDefault:"3000"`
		Topics         []string `env:"TOPICS" envSeparator:","`
	}

	Sasl struct {
		IsUse            bool   `env:"IS_USE" envDefault:"false"`
		Broker           string `env:"BROKER"`
		User             string `env:"USERNAME"`
		Password         string `env:"PASSWORD"`
		Mechanisms       string `env:"MECHANISMS"`
		SecurityProtocol string `env:"SECURITY_PROTOCOL"`
	}

	SchemaRegistry struct {
		Type string `env:"TYPE"`
		Url  string `env:"URL"`
	}

	Log struct {
		Level logger.LogLevel `env:"LEVEL" envDefault:"debug"`
	}
)

func New() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

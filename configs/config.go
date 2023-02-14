package configs

import (
	"github.com/caarlos0/env/v7"
)

type (
	Config struct {
		App            App            `envPrefix:"APP_"`
		Db             Db             `envPrefix:"DB_"`
		Kafka          Kafka          `envPrefix:"KAFKA_"`
		Sasl           Sasl           `envPrefix:"SASL_"`
		SchemaRegistry SchemaRegistry `envPrefix:"SR_"`
		Log            Log            `envPrefix:"LOG_"`
	}

	App struct {
		Name    string `env:"NAME,required"`
		Version string `env:"VERSION,required"`
	}

	Db struct {
		Connection string `env:"CONNECTION" envDefault:"pgsql"`
		Host       string `env:"HOST" envDefault:"localhost"`
		Port       int16  `env:"PORT" envDefault:"5432"`
		Name       string `env:"NAME,required"`
		User       string `env:"USER,required"`
		Password   string `env:"PASSWORD,required"`
	}

	Kafka struct {
		Broker         string `env:"BROKER,required"`
		SessionTimeout int32  `env:"SESSION_TIMEOUT" envDefault:"3000"`
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
		Level string `env:"LEVEL" envDefault:"debug"`
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

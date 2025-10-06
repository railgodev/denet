package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App        App
		HTTP       HTTP
		Log        Log
		Migrations Migrations
		PG         PG
		JWT       JWT
		Metrics    Metrics
		Swagger    Swagger
	}

	App struct {
		Name    string `env:"APP_NAME,required"`
		Version string `env:"APP_VERSION,required"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT,required"`
	}

	Log struct {
		Level string `env:"LOG_LEVEL,required"`
	}

	Migrations struct {
		MigratePath string `env:"MIGRATE_PATH,required"`
	}

	PG struct {
		URL string `env:"PG_URL,required"`
	}
	JWT struct {
		SECRET string `env:"JWT_SECRET,required"`
	}
	Metrics struct {
		Enabled bool `env:"METRICS_ENABLED" envDefault:"false"`
	}

	Swagger struct {
		Enabled bool `env:"SWAGGER_ENABLED" envDefault:"false"`
	}
)

func Load() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./.env", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

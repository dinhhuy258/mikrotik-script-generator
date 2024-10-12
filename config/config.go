package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	App struct {
		Name         string        `envconfig:"APP_NAME" default:"mikrotik-script-generator"`
		StartTimeout time.Duration `envconfig:"START_TIMEOUT" default:"1m"`
		StopTimeout  time.Duration `envconfig:"STOP_TIMEOUT" default:"1m"`
	}
	Http struct {
		Port         int           `envconfig:"HTTP_PORT" default:"8080"`
		ReadTimeout  time.Duration `envconfig:"HTTP_READ_TIMEOUT" default:"1m"`
		WriteTimeout time.Duration `envconfig:"HTTP_WRITE_TIMEOUT" default:"1m"`
	}
	Log struct {
		Level string `envconfig:"LOG_LEVEL" default:"debug"`
	}
}

func loadConfig() (*Config, error) {
	_ = godotenv.Load()

	var config Config

	err := envconfig.Process("mikrotik-script-generator", &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func NewConfig() *Config {
	config, err := loadConfig()
	if err != nil {
		log.Fatal("Error occurred when load config ", err)
	}

	return config
}

package config

import (
	env "github.com/caarlos0/env/v6"
)

type Config struct {
	MongoDb MongoDb
}

type MongoDb struct {
	HOST       string `env:"HOST"`
	PORT       string `env:"PORT"`
	DB         string `env:"DB"`
	COLLECTION string `env:"COLLECTION"`
}

func NewFromEnv() (*Config, error) {
	var config Config
	if err := env.Parse(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

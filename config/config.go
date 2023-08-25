package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	// add env tags to check which this struct is "Config" or not
	// what is "Env"…?
	Env  string `env:"TODO_ENV" envDefault:"dev"`
	Port int    `env:"PORT" envDefault:"80"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

package config

import (
	"github.com/vrischmann/envconfig"
)

type Config struct {
	Telegram   Telegram
	PostgreSQL PostgreSQL
}

type Telegram struct {
	Token string
}

type PostgreSQL struct {
	MasterDSN string
}

func GetConfig() (Config, error) {
	c := Config{}
	if err := envconfig.Init(&c); err != nil {
		return Config{}, err
	}

	return c, nil
}

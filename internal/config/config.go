package config

import (
	"github.com/vrischmann/envconfig"
)

type Config struct {
	EmWell     EmWell
	Telegram   Telegram
	PostgreSQL PostgreSQL
	Redis      Redis
}

type EmWell struct {
	URL string
}

type Telegram struct {
	Token               string
	LinkGeneratorSecret string
}

type PostgreSQL struct {
	MasterDSN string
}

type Redis struct {
	Host string
}

func GetConfig() (Config, error) {
	c := Config{}
	if err := envconfig.Init(&c); err != nil {
		return Config{}, err
	}

	return c, nil
}

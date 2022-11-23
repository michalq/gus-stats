package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

type Config struct {
	Gus struct {
		Client string `env:"GUS_CLIENT"`
	}
}

func LoadConfig() Config {
	var config Config
	if _, err := env.UnmarshalFromEnviron(&config); err != nil {
		log.Fatal(err)
	}
	return config
}

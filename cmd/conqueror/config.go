package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

func parseEnv() (*config, error) {
	c := new(config)
	err := envconfig.Process("", c)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse env")
	}

	return c, nil
}

type config struct {
	Port     string `envconfig:"PORT"`
	Secret   string `envconfig:"SECRET"`
	FilesDir string `envconfig:"FILES_DIR"`

	DBHost     string `envconfig:"DB_HOST"`
	DBName     string `envconfig:"DB_NAME"`
	DBUser     string `envconfig:"DB_USER"`
	DBPassword string `envconfig:"DB_PASSWORD"`

	MigrationsDir string `envconfig:"MIGRATIONS_DIR"`
}

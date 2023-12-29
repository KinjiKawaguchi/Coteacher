package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	DBUser string `envconfig:"MYSQL_USER" required:"true"`
	DBPass string `envconfig:"MYSQL_PASSWORD" required:"true"`
	DBAddr string `envconfig:"DB_ADDR" required:"true"`
	DBName string `envconfig:"MYSQL_DATABASE" required:"true"`
}

func New() (*Config, error) {
	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		return nil, err
	}
	return config, nil
}

package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Environment string
	AwsDynamodb DynamodbConfig `envconfig:"DYNAMODB"`
}

type DynamodbConfig struct {
	Endpoint string
}

func NewConfig() (*Config, error) {
	var conf Config
	if err := envconfig.Process("", &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

func (c *Config) IsLocal() bool {
	return c.Environment == "local"
}
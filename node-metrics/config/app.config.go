package config

import "os"

type AppConfig struct {
	serverPort string
}

func (c *AppConfig) GetServerAddress() string {
	return c.serverPort
}

func NewFromEnv() (*AppConfig, error) {
	return &AppConfig{
		serverPort: os.Getenv("PORT"),
	}, nil
}

package config

import "os"

type AppConfig struct {
	natsAddress     string
	serverAddress   string
	natsSubject     string
	magnetarAddress string
}

func (c *AppConfig) GetNatsAddress() string {
	return c.natsAddress
}
func (c *AppConfig) GetServerAddress() string {
	return c.serverAddress
}
func (c *AppConfig) GetMagnetarAddress() string { return c.magnetarAddress }
func (c *AppConfig) GetNatsSubject() string     { return c.natsSubject }

func NewFromEnv() (*AppConfig, error) {
	return &AppConfig{
		natsAddress:     os.Getenv("NATS_ADDRESS"),
		serverAddress:   os.Getenv("SERVER_ADDRESS"),
		magnetarAddress: os.Getenv("MAGNETAR_ADDRESS"),
		natsSubject:     os.Getenv("NATS_SUBJECT"),
	}, nil
}

package config

import (
	"os"
)

// Config ...
type Config struct {
	BindAddr    string
	DatabaseURL string
	SentryDSN   string
	Redis       Redis
	Payments    map[string]interface{}
}

// Redis ...
type Redis struct {
	Address  string
	Password string
	PoolSize int
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr:    os.Getenv("BIND_ADDR"),
		DatabaseURL: os.Getenv("DB_URL"),
		SentryDSN:   os.Getenv("SENTRY_DSN"),
		Redis: Redis{
			Address:  os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
			PoolSize: 2,
		},
	}
}

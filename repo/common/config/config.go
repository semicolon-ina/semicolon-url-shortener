package config

import (
	"time"
)

type DefaultConfig struct {
	// Tag envconfig langsung nembak nama variable di docker-compose
	AppName string `envconfig:"APP_NAME" default:"GoShort-Service"`
	AppPort string `envconfig:"APP_PORT" default:"8080"`
}

type URLConfig struct {
	BaseURL string `envconfig:"BASE_URL"`
}

type RedisConfig struct {
	// Nested struct di envconfig sebenernya otomatis nambah prefix (REDIS_HOST),
	// TAPI biar aman 100%, kita hardcode nama env-nya di tag.
	Enabled              bool          `envconfig:"REDIS_ENABLED"`
	Host                 string        `envconfig:"REDIS_HOST" default:"localhost"`
	Port                 string        `envconfig:"REDIS_PORT" default:"6379"`
	Username             string        `envconfig:"REDIS_USERNAME" default:"default"`
	Password             string        `envconfig:"REDIS_PASSWORD"`
	DB                   int           `envconfig:"REDIS_DB" default:"0"`
	EnableSSL            bool          `envconfig:"REDIS_ENABLE_SSL"`
	Expire               time.Duration `envconfig:"REDIS_EXPIRE" default:"1h"`
	ValidLatestTimestamp time.Duration `envconfig:"REDIS_VALID_LATEST_TIMESTAMP" default:"3h"`
}

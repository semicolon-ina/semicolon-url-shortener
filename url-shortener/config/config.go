package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/semicolon-ina/semicolon-url-shortener/repo/common/config"
)

type Config struct {
	config.DefaultConfig
	Redis config.RedisConfig
	URL   config.URLConfig
}

var cfg = Config{}

func LoadConfig() error {
	// 1. Coba load file .env (Khusus Local Development)
	// Di Docker file ini ga ada, jadi errornya kita ignore.
	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			fmt.Println("Warning: Error loading .env file")
		}
	}

	// 2. MAGIC-NYA DISINI:
	// envconfig akan scan semua Environment Variable (dari Docker Compose)
	// dan masukin ke struct berdasarkan tag `envconfig` yang kita buat tadi.
	// Prefix "" artinya dia baca langsung tanpa tambahan prefix global.
	if err := envconfig.Process("", &cfg); err != nil {
		return err
	}

	return nil
}

func Get() Config {
	return cfg
}

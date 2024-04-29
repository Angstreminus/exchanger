package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	ServerPort string
	ServerHost string
	LogLevel   string
}

// MustLoadConfig loads config from .env file or
func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
		return Config{}, err
	}
	once.Do(
		func() {
			config = &Config{
				ServerPort: os.Getenv("PORT"),
				ServerHost: os.Getenv("HOST"),
				LogLevel:   os.Getenv("LOG_LEVEL"),
			}
		})
	return *config, nil
}

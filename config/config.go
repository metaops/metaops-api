package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type DBConfig struct {
	Connection string
	Driver     string
}

type Config struct {
	DBConfig DBConfig
}

func Load() Config {

	return Config{
		DBConfig: DBConfig{
			Connection: os.Getenv("DB_URL"),
			Driver:     os.Getenv("DB_DRIVER"),
		},
	}
}

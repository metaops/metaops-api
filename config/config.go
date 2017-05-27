package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type DBConfig struct {
	Connection string
	Driver     string
}

type ServerConfig struct {
	Port string
}

type Config struct {
	DBConfig     DBConfig
	ServerConfig ServerConfig
}

func Load() Config {

	return Config{
		DBConfig: DBConfig{
			Connection: os.Getenv("DB_URL"),
			Driver:     os.Getenv("DB_DRIVER"),
		},
		ServerConfig: ServerConfig{
			Port: os.Getenv("PORT"),
		},
	}
}

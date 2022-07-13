package config

import (
	"os"
)

type ApiConfig struct {
	ApiHost string
	ApiPort string
}

type DbConfig struct {
	Host     string
	Port     string
	DbName   string
	User     string
	Password string
}

type Config struct {
	ApiConfig
	DbConfig
}

func (c Config) readConfig() Config {
	c.DbConfig = DbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DbName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	c.ApiConfig = ApiConfig{
		ApiHost: os.Getenv("API_HOST"),
		ApiPort: os.Getenv("API_PORT"),
	}
	return c
}

func NewConfig() Config {
	cfg := Config{}
	cfg.readConfig()
	return cfg
}

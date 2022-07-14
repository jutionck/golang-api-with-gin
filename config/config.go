package config

import (
	"os"
)

type FilePathConfig struct {
	FilePath string
}
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
	FilePathConfig
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
	c.FilePathConfig = FilePathConfig{FilePath: os.Getenv("FILE_PATH")}
	return c
}

func NewConfig() Config {
	cfg := Config{}
	return cfg.readConfig()
}

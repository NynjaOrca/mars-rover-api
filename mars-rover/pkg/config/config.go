package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	Postgres   Postgres
	NasaClient APIConfig
}

type Postgres struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     int
	SSLMode    string
}

type APIConfig struct {
	URL     string
	RelUrl  string
	Timeout int
	Key     string
}

func New() (Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Config{}, fmt.Errorf("Load ENV: %v", err)
	}

	var c Config
	c.Postgres, err = buildPostgresConfig()
	if err != nil {
		return Config{}, fmt.Errorf("POSTGRES Config: %v", err)
	}

	c.NasaClient, err = buildNasaConfig()
	if err != nil {
		return Config{}, fmt.Errorf("NASA Config: %v", err)
	}

	return c, nil
}

func buildPostgresConfig() (Postgres, error) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return Postgres{}, fmt.Errorf("DB_PORT: %v", err)
	}
	return Postgres{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		SSLMode:    os.Getenv("SSL_MODE"),
		DBPort:     port,
	}, nil
}

func buildNasaConfig() (APIConfig, error) {
	timeout, err := strconv.Atoi(os.Getenv("NASA_TIMEOUT"))
	if err != nil {
		return APIConfig{}, fmt.Errorf("NASA_TIMEOUT: %v", err)
	}
	return APIConfig{
		URL:     os.Getenv("NASA_URL"),
		RelUrl:  os.Getenv("NASA_RELATIVE"),
		Key:     os.Getenv("NASA_KEY"),
		Timeout: timeout,
	}, nil
}

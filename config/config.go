package config

import (
	"briefcash-consumer-bca/infrastructure/log"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBName     string
	DBUsername string
	DBPassword string
	AppPort    string
	RedisHost  string
	RedisPort  string
	KafkaHost  string
	KafkaPort  string
	KafkaTopic string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Logger.WithError(err).Warn("Failed to load credentials from .env file, checking credentials in OS environement variable")
	}

	cfg := &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		RedisHost:  os.Getenv("REDIS_HOST"),
		RedisPort:  os.Getenv("REDIS_PORT"),
		KafkaHost:  os.Getenv("KAFKA_HOST"),
		KafkaPort:  os.Getenv("KAFKA_PORT"),
		KafkaTopic: os.Getenv("KAFKA_TOPIC"),
		AppPort: func() string {
			if value := os.Getenv("APP_PORT"); value != "" {
				return value
			}
			return ":8080"
		}(),
	}

	if cfg.DBHost == "" {
		log.Logger.Error("DB_HOST not found in environement")
		return nil, fmt.Errorf("DB_HOST not found in environment")
	}

	if cfg.RedisHost == "" {
		log.Logger.Error("REDIS_HOST not found in environement")
		return nil, fmt.Errorf("REDIS_HOST not found in environment")
	}

	if cfg.KafkaHost == "" {
		log.Logger.Error("KAFKA_HOST not found in environement")
		return nil, fmt.Errorf("KAFKA_HOST not found in environment")
	}

	return cfg, nil
}

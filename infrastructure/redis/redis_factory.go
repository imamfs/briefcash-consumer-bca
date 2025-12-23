package redis

import (
	"briefcash-consumer-bca/config"
	"briefcash-consumer-bca/infrastructure/log"
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisFactory struct {
	Client *redis.Client
}

func NewRedisFactory(cfg *config.Config) (*RedisFactory, error) {
	if cfg.RedisHost == "" || cfg.RedisPort == "" {
		log.Logger.Error("Redis port or host is not configured in environement")
		return nil, fmt.Errorf("value in redis port or host is empty")
	}

	address := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

	client := redis.NewClient(&redis.Options{
		Addr:         address,
		DB:           0,
		PoolSize:     50,
		MinIdleConns: 10,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		log.Logger.WithError(err).Error("Failed to ping redis server")
		return nil, fmt.Errorf("failed to ping redis server, with error: %w", err)
	}

	log.Logger.Info("Redis server successfully connected!")

	return &RedisFactory{Client: client}, nil
}

func (r *RedisFactory) Close() error {
	return r.Client.Close()
}

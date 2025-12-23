package database

import (
	"briefcash-consumer-bca/config"
	"context"
	"fmt"
	"time"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbAdapter struct {
	DB *gorm.DB
}

func NewDbAdapter(cfg config.Config) (*DbAdapter, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUsername, cfg.DBPassword, cfg.DBName, "disable",
	)

	psql, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: false,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("failed to establish database connection, with error: %w", err)
	}

	sql, err := psql.DB()

	if err != nil {
		return nil, fmt.Errorf("failed to get generic database, with error: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sql.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database, with error: %w", err)
	}

	sql.SetMaxIdleConns(10)
	sql.SetMaxOpenConns(100)
	sql.SetConnMaxLifetime(time.Hour)

	return &DbAdapter{DB: psql}, nil
}

func (adapter *DbAdapter) Close() error {
	sql, err := adapter.DB.DB()

	if err == nil {
		return sql.Close()
	}

	return fmt.Errorf("failed to close database connection, with error: %w", err)
}

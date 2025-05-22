package main

import (
	"API_Service/internal/config"
	"API_Service/internal/storage/sqlite"
	"go.uber.org/zap"
)

const (
	envProd = "prod"
	envDev  = "dev"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting service", zap.String("config", cfg.Env))
	log.Debug("debug message are enable")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Fatal("failed to init storage", zap.Error(err))
	}

	_ = storage
}

func setupLogger(env string) *zap.Logger {
	var log *zap.Logger
	switch env {
	case envDev:
		log, _ = zap.NewDevelopment()
	case envProd:
		log, _ = zap.NewProduction()
	}
	return log
}

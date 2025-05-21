package main

import (
	"API_Service/internal/config"
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

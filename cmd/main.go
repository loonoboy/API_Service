package main

import (
	"API_Service"
	"API_Service/internal/config"
	"API_Service/internal/handler"
	"API_Service/internal/repository"
	"API_Service/internal/repository/postgres"
	"API_Service/internal/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
	"os"
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

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file", zap.Error(err))
	}

	db, err := postgres.NewPostgresDB(config.DB{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  cfg.DB.SSLMode,
	})
	if err != nil {
		log.Fatal("failed to connect to database", zap.Error(err))
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(log, services)

	log.Info("starting server", zap.String("address", cfg.Addr))

	srv := new(API_Service.Server)
	if err = srv.Run(config.HTTPServer{
		Addr:        cfg.HTTPServer.Addr,
		Timeout:     cfg.HTTPServer.Timeout,
		IdleTimeout: cfg.HTTPServer.IdleTimeout,
	}, handlers.InitRoutes()); err != nil {
		log.Fatal("failed to start server", zap.Error(err))
	}
	log.Error("server stopped")
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

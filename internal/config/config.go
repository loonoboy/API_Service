package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
	DB          `yaml:"db"`
}

type HTTPServer struct {
	Addr        string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type DB struct {
	Host     string `yaml:"host" env-required:"true"`
	Port     string `yaml:"port" env-required:"true"`
	SSLMode  string `yaml:"ssl_mode" env-required:"true"`
	Username string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_Password"`
	DBName   string `env:"POSTGRES_DB"`
}

func MustLoad() *Config {
	//mb use viper
	err := os.Setenv("CONFIG_PATH", "./config/local.yaml")
	if err != nil {
		log.Fatal("can't load config path", err)
	}
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH environment variable not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("CONFIG_PATH does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}

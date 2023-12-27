package config

import (
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	AuthAccessTelegramToken       string `env:"AUTH_ACCESS_TELEGRAM_TOKEN,required"`
	AuthAccessOpenweathermapToken string `env:"AUTH_ACCESS_OPENWEATHERMAP_TOKEN,required"`
	Environment                   string `env:"ENVIRONMENT,required"`
	LogLevel                      string `env:"LOG_LEVEL,required"`
	MongoDbUri                    string `env:"MONGODB_URI,required"`
	MongoDbUser                   string `env:"MONGODB_USER,required"`
	MongoDbPass                   string `env:"MONGODB_PASS,required"`
	MongoDbName                   string `env:"MONGODB_NAME,required"`
}

func CreateConfig() *Config {
	return &Config{}
}

func LoadEnvConfig(envStr string) (*Config, error) {
	err := godotenv.Load(envStr)
	if err != nil {
		return nil, err
	}

	cfg := CreateConfig()
	err = env.Parse(cfg)
	if err != nil {
		return cfg, err
	}

	logrus.SetFormatter(&logrus.JSONFormatter{})
	if cfg.Environment == "development" {
		logrus.SetFormatter(&logrus.TextFormatter{})
		logrus.SetOutput(os.Stdout)
	}
	level, err := logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		return cfg, err
	}
	logrus.SetLevel(level)

	return cfg, nil
}

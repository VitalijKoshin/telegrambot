package main

import (
	"telegrambot/database/mongodb"
	"telegrambot/handlers"
	"telegrambot/internal/config"
	"telegrambot/pkg/telegramapi"
	"telegrambot/pkg/weatherapi"
	repository "telegrambot/repository/mongodb"
	"telegrambot/services"

	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.LoadEnvConfig(".env")
	if err != nil {
		logrus.Panic(err)
	}
	mdbClient, err := mongodb.NewClient(cfg)
	if err != nil {
		logrus.Panic(err)
	}
	defer mdbClient.Disconnect()
	userRepo := repository.NewUserRepository(cfg, *mdbClient)
	subsRepo := repository.NewSubsRepository(cfg, *mdbClient)
	uService := services.NewUserService(userRepo)
	sService := services.NewSubsService(subsRepo)

	weatherClient := weatherapi.CreateClient(cfg.AuthAccessOpenweathermapToken)
	bot, err := telegramapi.CreateBot(cfg.AuthAccessTelegramToken)
	if err != nil {
		logrus.Panic(err)
	}
	tService := services.NewTicker(10, uService, sService, bot, weatherClient)
	tService.StartTicker()
	messageHandler := handlers.CreateMessageHanlder(weatherClient, uService, sService, bot)
	messageHandler.ListenBot()
}

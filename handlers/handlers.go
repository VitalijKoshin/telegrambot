package handlers

import (
	"telegrambot/pkg/telegramapi"
	"telegrambot/pkg/utils"
	weather "telegrambot/pkg/weatherapi"
	"telegrambot/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

const (
	typeMsgIgnor    = -1
	typeMsgText     = 0
	typeMsgCommand  = 1
	typeMsgCallBack = 2
)

type MessageHandler struct {
	weatherClient weather.WeatherClient
	userService   services.UserService
	subsService   services.SubsService
	tBot          telegramapi.Telegabot
}

func CreateMessageHanlder(wc weather.WeatherClient, us services.UserService, ss services.SubsService, tb telegramapi.Telegabot) *MessageHandler {
	return &MessageHandler{weatherClient: wc, userService: us, subsService: ss, tBot: tb}
}

func (mh *MessageHandler) ListenBot() {
	mh.InitCommandMap()
	mh.InitCallbackMap()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	bot := mh.tBot.GetBot()
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		typeMsg := getTypeMsg(&update)
		switch typeMsg {
		case typeMsgCommand:
			mh.HandleCommand(bot, update.Message)
		case typeMsgCallBack:
			mh.HandleCallbackQuery(bot, update.CallbackQuery)
		case typeMsgText:
			mh.HandleMessage(bot, update.Message)
		default:
			continue
		}
	}
}

func getTypeMsg(up *tgbotapi.Update) int {
	if up.CallbackQuery != nil {
		return typeMsgCallBack
	}
	if up.Message == nil {
		return typeMsgIgnor
	}
	if up.Message.IsCommand() {
		return typeMsgCommand
	}
	return typeMsgText
}

func (mh *MessageHandler) handleForecast(msg string) (string, []tgbotapi.MessageEntity) {
	geo, err := utils.PrepareCoordinats(msg)
	if err != nil {
		logrus.Debug(err)
		return err.Error(), nil
	}

	geoWheather := &weather.GeoLocation{Latitude: geo.Latitude, Longitude: geo.Longitude, Exclude: "current"}
	forecast, err := mh.weatherClient.GetWeatherForecastByLocation(*geoWheather)
	if err != nil {
		logrus.Debug(err)
		return "Somthing went wrong. Try again.", nil
	}
	return forecast, utils.FormatForecast(forecast)
}

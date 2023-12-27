package services

import (
	"fmt"
	"time"

	"telegrambot/models"
	"telegrambot/pkg/telegramapi"
	"telegrambot/pkg/utils"
	"telegrambot/pkg/weatherapi"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type ServiceTickerI interface {
	StartTicker()
}
type ServiceTicker struct {
	wc          weatherapi.WeatherClient
	userService UserService
	subsService SubsService
	tBot        telegramapi.Telegabot
	tgBot       *tgbotapi.BotAPI
	ticker      *time.Ticker
	done        chan bool
	seconds     time.Duration
}

func NewTicker(secondsPeriod time.Duration, us UserService, ss SubsService, tb telegramapi.Telegabot, wc weatherapi.WeatherClient) ServiceTickerI {
	return &ServiceTicker{
		userService: us,
		subsService: ss,
		tBot:        tb,
		tgBot:       tb.GetBot(),
		wc:          wc,
		ticker:      time.NewTicker(secondsPeriod * time.Second),
		seconds:     secondsPeriod,
		done:        make(chan bool),
	}
}

func (st *ServiceTicker) StartTicker() {
	go func() {
		for {
			select {
			case <-st.done:
				return
			case t := <-st.ticker.C:
				fmt.Println("Tick at", t)
				st.sendForecastByRange()
			}
		}
	}()
}

func (st *ServiceTicker) sendForecastByRange() {
	currentTS := st.subsService.GetUTCTimeStamp64()
	ss, err := st.subsService.FindSubsBetweenTime((currentTS - 86400), currentTS)
	if err != nil {
		return
	}
	for _, s := range *ss {
		_ = st.sendBotMsg(s)
		s.LTime = st.subsService.GetUTCTimeStamp64()
		s.NTime = s.LTime + s.Frequency
		st.subsService.SaveSubs(&s)
	}
}

func (st *ServiceTicker) sendBotMsg(s models.Subs) error {
	geoWheather := weatherapi.GeoLocation{Latitude: s.Latitude, Longitude: s.Longitude, Exclude: "current"}
	forecast, err := st.wc.GetWeatherForecastByLocation(geoWheather)
	if err != nil {
		return err
	}
	entities := utils.FormatForecast(forecast)
	response := tgbotapi.NewMessage(s.ChatID, forecast)
	response.Entities = entities
	logrus.Debug(response.Text)
	if _, err := st.tgBot.Send(response); err != nil {
		logrus.Panic(err)
		return err
	}
	return nil
}

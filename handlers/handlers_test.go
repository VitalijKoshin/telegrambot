package handlers

import (
	"testing"

	"telegrambot/pkg/utils"
	weather "telegrambot/pkg/weatherapi"
	"telegrambot/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

func Test_getTypeMsg(t *testing.T) {
	type args struct {
		up *tgbotapi.Update
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Ignor", args{up: &tgbotapi.Update{}}, typeMsgIgnor},
		{"CallbackQuery", args{up: &tgbotapi.Update{CallbackQuery: &tgbotapi.CallbackQuery{}}}, typeMsgCallBack},
		{"Message", args{up: &tgbotapi.Update{Message: &tgbotapi.Message{}}}, typeMsgText},
		{"Command", args{
			up: &tgbotapi.Update{
				Message: &tgbotapi.Message{Entities: []tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0}}},
			},
		}, typeMsgCommand},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTypeMsg(tt.args.up)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestMessageHandler_handleForecast(t *testing.T) {
	type fields struct {
		weatherClient weather.WeatherClient
		userService   services.UserService
		subsService   services.SubsService
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  []tgbotapi.MessageEntity
	}{
		{
			name: "Usual",
			fields: fields{
				weatherClient: &weatherClientImplFake{},
				userService:   &UserServiceFake{},
				subsService:   &SubsServiceFake{},
			},
			args:  args{msg: "90:90"},
			want:  "GetWeatherForecastByLocation",
			want1: []tgbotapi.MessageEntity{},
		},
		{
			name: "Error",
			fields: fields{
				weatherClient: &weatherClientImplFakeError{},
				userService:   &UserServiceFake{},
				subsService:   &SubsServiceFake{},
			},
			args:  args{msg: "90:90"},
			want:  "Somthing went wrong. Try again.",
			want1: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mh := &MessageHandler{
				weatherClient: tt.fields.weatherClient,
				userService:   tt.fields.userService,
				subsService:   tt.fields.subsService,
			}
			got, got1 := mh.handleForecast(tt.args.msg)
			tt.want1 = utils.FormatForecast(tt.want)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, got1, tt.want1)
		})
	}
}

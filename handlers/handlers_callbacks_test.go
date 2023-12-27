package handlers

import (
	"testing"

	"telegrambot/models"
	weather "telegrambot/pkg/weatherapi"
	"telegrambot/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
)

func TestMessageHandler_callbackSubsForecastFunc(t *testing.T) {
	type fields struct {
		weatherClient weather.WeatherClient
		userService   services.UserService
		subsService   services.SubsService
	}
	type args struct {
		cds callbackDataStruct
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
		want1  []tgbotapi.MessageEntity
	}{
		{
			name: "Empty args",
			fields: fields{
				weatherClient: &weatherClientImplFake{},
				userService:   &UserServiceFake{},
				subsService:   &SubsServiceFake{},
			},
			args: args{
				cds: callbackDataStruct{callbackStr: "", callbackArg: ""},
			},
			want:  "",
			want1: []tgbotapi.MessageEntity{},
		},
		{
			name: "Empty args callbackStr",
			fields: fields{
				weatherClient: &weatherClientImplFake{},
				userService:   &UserServiceFake{},
				subsService:   &SubsServiceFake{},
			},
			args: args{
				cds: callbackDataStruct{
					callbackStr: "",
					callbackArg: "30:90",
					user:        &models.User{},
					chatID:      777,
				},
			},
			want:  "subscribed",
			want1: nil,
		},
		{
			name: "Empty args callbackArg not exist callback",
			fields: fields{
				weatherClient: &weatherClientImplFake{},
				userService:   &UserServiceFake{},
				subsService:   &SubsServiceFake{},
			},
			args: args{
				cds: callbackDataStruct{
					callbackStr: "wefcwef",
					callbackArg: "",
					user:        &models.User{},
					chatID:      777,
				},
			},
			want:  "",
			want1: []tgbotapi.MessageEntity{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mh := &MessageHandler{
				weatherClient: &weatherClientImplFake{},
				userService:   &UserServiceFake{},
				subsService:   &SubsServiceFake{},
			}
			got, got1 := mh.callbackSubsForecastFunc(tt.args.cds)
			assert.Equal(t, got, tt.want)
			assert.Equal(t, got1, tt.want1)
		})
	}
}

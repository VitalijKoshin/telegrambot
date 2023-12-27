package telegramapi

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegabot interface {
	GetBot() *tgbotapi.BotAPI
}

type telegabotImpl struct {
	Bot *tgbotapi.BotAPI
}

func CreateBot(authAccessTelegramToken string) (Telegabot, error) {
	bot, err := tgbotapi.NewBotAPI(authAccessTelegramToken)
	if err != nil {
		return nil, err
	}
	bot.Debug = true
	return &telegabotImpl{Bot: bot}, nil
}

func (t telegabotImpl) GetBot() *tgbotapi.BotAPI {
	return t.Bot
}

package handlers

import (
	"telegrambot/pkg/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func (mh *MessageHandler) HandleMessage(bot *tgbotapi.BotAPI, tmsg *tgbotapi.Message) {
	txt, entities := mh.handleText(tmsg.Text)
	response := tgbotapi.NewMessage(tmsg.Chat.ID, txt)
	response.Entities = entities
	if utils.TextHasCoordinats(tmsg.Text) {
		response.ReplyMarkup = mh.HandleKeyboardButton(tmsg.Text, tmsg)
	}
	logrus.Debug(response.Text)
	if _, err := bot.Send(response); err != nil {
		logrus.Panic(err)
	}
}

func (mh *MessageHandler) handleText(msg string) (string, []tgbotapi.MessageEntity) {
	errStr, result := mh.handleForecast(msg)
	if result != nil {
		return errStr, result
	}
	return errStr, result
}

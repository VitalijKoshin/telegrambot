package handlers

import (
	"fmt"
	"strings"
	"time"

	"telegrambot/models"
	"telegrambot/pkg/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

var callbacksMap map[string]func(callbackDataStruct) (string, []tgbotapi.MessageEntity)

type callbackDataStruct struct {
	callbackStr string
	callbackArg string
	user        *models.User
	chatID      int64
}

func (mh *MessageHandler) InitCallbackMap() {
	callbacksMap = map[string]func(callbackDataStruct) (string, []tgbotapi.MessageEntity){
		"callback_subs_forecast":   mh.callbackSubsForecastFunc,
		"callback_unsubs_forecast": mh.callbackUnSubsForecastFunc,
	}
}

func (mh *MessageHandler) HandleCallbackQuery(bot *tgbotapi.BotAPI, callbackQuery *tgbotapi.CallbackQuery) {
	callback := tgbotapi.NewCallback(callbackQuery.ID, callbackQuery.Data)
	if _, err := bot.Request(callback); err != nil {
		logrus.Panic(err)
	}
	cds, err := createCallbackDataStruct(callbackQuery.Data)
	if err != nil {
		logrus.Debug(err)
		return
	}
	cds.user, err = mh.userService.LoadUserByUsername(callbackQuery.From.UserName)
	if err != nil {
		logrus.Debug(err)
		return
	}
	cds.chatID = callbackQuery.Message.Chat.ID
	if callbackFunc, existCommand := callbacksMap[cds.callbackStr]; existCommand {
		callbackQuery.Data, _ = callbackFunc(cds)
	}

	msg := tgbotapi.NewMessage(callbackQuery.Message.Chat.ID, callbackQuery.Data)
	if _, err := bot.Send(msg); err != nil {
		logrus.Panic(err)
	}
}

func (mh *MessageHandler) callbackSubsForecastFunc(cds callbackDataStruct) (string, []tgbotapi.MessageEntity) {
	geo, err := utils.PrepareCoordinats(cds.callbackArg)
	if err != nil {
		logrus.Debug(err)
		return "", []tgbotapi.MessageEntity{}
	}
	cTime := time.Time.Unix(time.Time{})
	ss := mh.subsService.CreateSubs(cds.chatID, cTime, cds.user.ID, geo.Latitude, geo.Longitude, 86400)
	_, err = mh.subsService.SaveSubs(&ss)
	if err != nil {
		logrus.Debug(err)
		return "", []tgbotapi.MessageEntity{}
	}
	return "subscribed", nil
}

func (mh *MessageHandler) callbackUnSubsForecastFunc(cds callbackDataStruct) (string, []tgbotapi.MessageEntity) {
	geo, err := utils.PrepareCoordinats(cds.callbackArg)
	if err != nil {
		return "", []tgbotapi.MessageEntity{}
	}
	cTime := time.Time.Unix(time.Time{})
	ss := mh.subsService.CreateSubs(cds.chatID, cTime, cds.user.ID, geo.Latitude, geo.Longitude, 86400)
	s, err := mh.subsService.FindSubsExist(&ss)
	if err != nil {
		logrus.Debug(err)
		return "unsubscribe error", []tgbotapi.MessageEntity{}
	}
	if s == nil {
		return "subscribe is not exist", []tgbotapi.MessageEntity{}
	}
	_, err = mh.subsService.DeleteSubs(s)
	if err != nil {
		logrus.Debug(err)
		return "unsubscribe error", []tgbotapi.MessageEntity{}
	}
	return "unsubscribed", []tgbotapi.MessageEntity{}
}

func createCallbackDataStruct(callbackData string) (callbackDataStruct, error) {
	var cds callbackDataStruct
	data := strings.Split(callbackData, ":")
	if len(data) == 1 {
		return cds, fmt.Errorf("arg is not exist")
	}
	cds.callbackStr = data[0]
	cds.callbackArg = strings.Replace(callbackData, data[0]+":", "", 1)
	return cds, nil
}

func (mh *MessageHandler) HandleKeyboardButton(geoStr string, tmsg *tgbotapi.Message) tgbotapi.InlineKeyboardMarkup {
	user, err := mh.userService.LoadUserByUsername(tmsg.From.UserName)
	if err != nil {
		logrus.Debug(err)
		return tgbotapi.InlineKeyboardMarkup{}
	}
	geo, err := utils.PrepareCoordinats(geoStr)
	if err != nil {
		logrus.Debug(err)
		return tgbotapi.InlineKeyboardMarkup{}
	}
	cTime := time.Time.Unix(time.Time{})
	ss := mh.subsService.CreateSubs(tmsg.Chat.ID, cTime, user.ID, geo.Latitude, geo.Longitude, 86400)
	s, err := mh.subsService.FindSubsExist(&ss)
	if err != nil {
		logrus.Debug(err)
		return tgbotapi.InlineKeyboardMarkup{}
	}
	if s == nil {
		return mh.createInlineKeybordSubscribe(geoStr)
	}
	return mh.createInlineKeybordUnSubscribe(geoStr)
}

func (mh *MessageHandler) createInlineKeybordSubscribe(geoStr string) tgbotapi.InlineKeyboardMarkup {
	callBackStr := fmt.Sprintf("callback_subs_forecast:%s", geoStr)
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("subscribe", callBackStr),
		),
	)
}

func (mh *MessageHandler) createInlineKeybordUnSubscribe(geoStr string) tgbotapi.InlineKeyboardMarkup {
	callBackStr := fmt.Sprintf("callback_unsubs_forecast:%s", geoStr)
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("unsubscribe", callBackStr),
		),
	)
}

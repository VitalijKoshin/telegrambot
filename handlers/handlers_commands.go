package handlers

import (
	"fmt"

	"telegrambot/pkg/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

const defaultCommandName = "default"

var commandsMap map[string]func(*commandDataStruct) (string, []tgbotapi.MessageEntity)

type commandDataStruct struct {
	tMsg *tgbotapi.Message
}

func (mh *MessageHandler) InitCommandMap() {
	commandsMap = map[string]func(*commandDataStruct) (string, []tgbotapi.MessageEntity){
		"forecastlist":     mh.botsHandlerForecastList,
		"forecast":         mh.botsHandlerForecast,
		"help":             mh.botsHandlerHelp,
		"start":            mh.botsHandlerStart,
		defaultCommandName: mh.botsHandlerDefault,
	}
}

func (mh *MessageHandler) HandleCommand(bot *tgbotapi.BotAPI, tmsg *tgbotapi.Message) {
	txt, entities := mh.callCommand(tmsg.Command(), tmsg)
	response := tgbotapi.NewMessage(tmsg.Chat.ID, txt)

	if utils.CommandHasCoordinats(tmsg.Command(), tmsg.Text) {
		geoStr := utils.CommandGetStrCoordinats(tmsg.Command(), tmsg.Text)
		response.ReplyMarkup = mh.HandleKeyboardButton(geoStr, tmsg)
	}
	response.Entities = entities
	logrus.Debug(response.Text)
	if _, err := bot.Send(response); err != nil {
		logrus.Panic(err)
	}
}

func (mh *MessageHandler) callCommand(nameCommand string, tmsg *tgbotapi.Message) (string, []tgbotapi.MessageEntity) {
	cds := &commandDataStruct{tMsg: tmsg}
	if commandFunc, existCommand := commandsMap[nameCommand]; existCommand {
		return commandFunc(cds)
	}

	return commandsMap[defaultCommandName](cds)
}

func (mh *MessageHandler) botsHandlerForecastList(cd *commandDataStruct) (string, []tgbotapi.MessageEntity) {
	user, err := mh.userService.LoadUserByUsername(cd.tMsg.From.UserName)
	if err != nil {
		logrus.Debug(err.Error())
		return err.Error(), nil
	}
	flist, err := mh.subsService.FindSubsByUID(user.ID)
	if err != nil {
		logrus.Debug(err.Error())
		return err.Error(), nil
	}
	result := "Forecasts List:\n"
	for _, v := range *flist {
		tmpStr := fmt.Sprintf("%f:%f\n", v.Latitude, v.Longitude)
		result += tmpStr
	}
	return result, nil
}

func (mh *MessageHandler) botsHandlerForecast(cd *commandDataStruct) (string, []tgbotapi.MessageEntity) {
	geoStr := utils.CommandGetStrCoordinats(cd.tMsg.Command(), cd.tMsg.Text)
	txt, result := mh.handleForecast(geoStr)
	if result != nil {
		return txt, result
	}

	return txt, result
}

func (mh *MessageHandler) botsHandlerDefault(cd *commandDataStruct) (string, []tgbotapi.MessageEntity) {
	return "I don't know that command", nil
}

func (mh *MessageHandler) botsHandlerStart(cd *commandDataStruct) (string, []tgbotapi.MessageEntity) {
	uc := mh.userService.CreateUser(cd.tMsg.From.ID, cd.tMsg.From.UserName, cd.tMsg.From.FirstName, cd.tMsg.From.LastName)
	u, err := mh.userService.AddUser(&uc)
	if err != nil {
		logrus.Debug(err)
		return err.Error(), nil
	}

	output := fmt.Sprintf("Welcome %s to the wheather forecast bot \n"+
		"/help - shows a list of commands\n", u.FirstName)

	return output, nil
}

func (mh *MessageHandler) botsHandlerHelp(cd *commandDataStruct) (string, []tgbotapi.MessageEntity) {
	result := "/start - welcome to the wheather forecast bot \n" +
		"/help - shows a list of commands\n" +
		"longitude:latitude - view forecast by geo longitude:latitude\n" +
		"/forecast longitude:latitude - view forecast by geo longitude:latitude\n" +
		"/forecastlist - your list of the forecast's subscriptions\n"

	return result, nil
}

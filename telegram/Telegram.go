package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
	"reflect"
	"runtime"

	"github.com/NekitMalyarenko/VocabularyBot/types"
	"github.com/NekitMalyarenko/VocabularyBot/bot/services"
)


type Bot struct {
	*tgbotapi.BotAPI
}


const (
	ME = 360952996
)


func Init(botToken string, debug bool) *Bot {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	bot.Debug = debug
	bot.Send(tgbotapi.NewMessage(ME, "alive"))

	return &Bot {
		bot,
	}
}


func (bot *Bot) Start() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		handleUpdate(bot, update)
	}
	
	return nil
}


func handleUpdate(bot *Bot, update tgbotapi.Update) {
	//sec, nan := time.Now().Second(), time.Now().Nanosecond()

	var (
		actionData = new(types.ActionData)
		action       func(data *types.ActionData) bool
	)

	actionData.ChatId = getChatId(update)

	if !getContextHolder().hasContext(actionData.ChatId) {
		getContextHolder().createContext(actionData.ChatId)
	}

	actionData.Bot     = bot.BotAPI
	actionData.Update  = update
	actionData.Context = getContextHolder().getContext(actionData.ChatId)
	
	if update.CallbackQuery == nil {
		if actionData.Context.NextAction != nil {
			action = actionData.Context.NextAction
		} else {
			action = getTextActionsHolder().getAction(strings.ToLower(update.Message.Text))
		}
	} else {
		if actionData.Context.NextAction != nil {
			actionData.Context.Clear()
		}

		buttonData, actionId, err := telegramServices.ButtonInit(update.CallbackQuery.Data)
		if err != nil {
			log.Println(err)
			bot.sendErrorMessage(actionData.ChatId)
			return
		}
		actionData.ButtonData = buttonData
		action = getButtonActionsHolder().getAction(actionId)
	}


	if action == nil {
		bot.sendActionNotFound(actionData.ChatId)
	} else {
		if !action(actionData) {
			bot.sendErrorMessage(actionData.ChatId)
		}
	}

	//log.Println("TIME:", time.Now().Second() - sec, ".", time.Now().Nanosecond() - nan)
}


func (bot *Bot) RegisterTextHandler(key string, function func(*types.ActionData)bool) {
	getTextActionsHolder().registerAction(key, function)
}


func (bot *Bot) RegisterButtonHandler(function func(*types.ActionData)bool) {
	getButtonActionsHolder().registerAction(function)
}


func GetFuncId(input func(*types.ActionData)bool) string {
	return runtime.FuncForPC(reflect.ValueOf(input).Pointer()).Name()
}


func (bot *Bot) sendActionNotFound(chatId int64) {
	message := tgbotapi.NewMessage(chatId, "Sorry, but I don't understand you =(")
	bot.Send(message)
}


func (bot *Bot) sendErrorMessage(chatId int64) {
	message := tgbotapi.NewMessage(chatId, "Oppss... Something went wrong =(")
	bot.Send(message)
}


func getChatId(update tgbotapi.Update) int64 {

	if update.CallbackQuery == nil {
		return update.Message.Chat.ID
	} else {
		return update.CallbackQuery.Message.Chat.ID
	}

}
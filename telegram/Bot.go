package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"vars"
	"telegram/data"
	"telegram/handlers"
)

var (
	bot *tgbotapi.BotAPI
)


func Start() {
	actionsInit()
	botInit()
}


func botInit() {
	var err error

	bot, err = tgbotapi.NewBotAPI(vars.GetString(vars.TELEGRAM_BOT_TOKEN))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		handleUpdate(update)
	}
}


func handleUpdate(update tgbotapi.Update) {
	var (
		context *telegramData.Context
		chatId = getChatId(update)
		action func(data telegramData.ActionData)bool
	)

	if !telegramData.GetContextHolder().HasContext(update.Message.Chat.ID) {
		telegramData.GetContextHolder().CreateContext(chatId)
	}
	context = telegramData.GetContextHolder().GetContext(chatId)

	if context.NextAction == nil {

		if update.CallbackQuery == nil {
			action = telegramData.GetActionsHolder().GetAction(update.Message.Text)
		}

	} else {
		action = context.NextAction
	}

	if action == nil {
		sendActionNotFound(update)
	} else {
		data := telegramData.ActionData{
			Update:  update,
			Context: telegramData.GetContextHolder().GetContext(update.Message.Chat.ID),
			ChatId:  update.Message.Chat.ID,
			Bot:     bot,
		}
		if !action(data) {
			sendErrorMessage(update)
		}
	}

}


func sendActionNotFound(update tgbotapi.Update) {
	message := tgbotapi.NewMessage(getChatId(update), "Sorry, but I don't understand you =(")
	bot.Send(message)
}


func sendErrorMessage(update tgbotapi.Update) {
	message := tgbotapi.NewMessage(getChatId(update), "Oppss... Something went wrong =(")
	bot.Send(message)
}


func getChatId(update tgbotapi.Update) int64 {

	if update.CallbackQuery == nil {
		return update.Message.Chat.ID
	} else {
		return update.CallbackQuery.Message.Chat.ID
	}

}


func actionsInit() {
	telegramData.GetActionsHolder().AddAction("/start", telegramHandlers.NewUserActionHandler)
	telegramData.GetActionsHolder().AddAction("/begin_test", telegramHandlers.BeginNNTraining)
	telegramData.GetActionsHolder().AddAction("/getWord", telegramHandlers.GetWord)
	//telegramData.GetActionsHolder().AddAction("/cancel", telegramHandlers.Cancel)
}
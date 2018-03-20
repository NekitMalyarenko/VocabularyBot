package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"

	"github.com/NekitMalyarenko/VocabularyBot/telegram/data"
	"github.com/NekitMalyarenko/VocabularyBot/telegram/handlers"
	"github.com/NekitMalyarenko/VocabularyBot/vars"
	"github.com/NekitMalyarenko/VocabularyBot/telegram/helpers"
)


var (
	bot *tgbotapi.BotAPI
)


func Start() {
	textActionsInit()
	buttonActionsInit()

	if vars.GetBoolean(vars.BOT_LEARNING) {
		log.Println("BOT LEARNING IS TRUE")
		go initBotLearning()
	}

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
		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		handleUpdate(update)
	}
}


func textActionsInit() {
	telegramData.GetActionsHolder().AddAction("/start", telegramHandlers.NewUserActionHandler)
	telegramData.GetActionsHolder().AddAction("/begin_test", telegramHandlers.BeginNNTraining)
	telegramData.GetActionsHolder().AddAction("/getWord", telegramHandlers.GetWord)
	//telegramData.GetActionsHolder().AddAction("/cancel", telegramHandlers.Cancel)
}


func buttonActionsInit() {
	telegramData.GetButtonsHolder().RegisterButton(telegramHandlers.START_NN_TRAINIG_BUTTON, telegramHandlers.BeginNNTrainingButton)
}


func handleUpdate(update tgbotapi.Update) {
	var (
		context    *telegramData.Context
		buttonData map[string]interface{}
		action     func(data telegramData.ActionData) bool
		err        error

		chatId = getChatId(update)
	)

	if !telegramData.GetContextHolder().HasContext(chatId) {
		telegramData.GetContextHolder().CreateContext(chatId)
	}
	context = telegramData.GetContextHolder().GetContext(chatId)


	if update.CallbackQuery == nil {
		if context.NextAction != nil {
			action = context.NextAction
		} else {
			action = telegramData.GetActionsHolder().GetAction(update.Message.Text)
		}
	} else {
		if context.NextAction != nil {
			context.ReCreateContext()
		}

		buttonData, action, err = telegramHelpers.ButtonInit(update.CallbackQuery.Data)
		if err != nil {
			log.Println(err)
			sendErrorMessage(update)
			return
		}
	}


	if action == nil {
		sendActionNotFound(update)
	} else {
		data := telegramData.ActionData{
			Update     :  update,
			Context    : telegramData.GetContextHolder().GetContext(chatId),
			ButtonData : buttonData,
			ChatId     :  chatId,
			Bot        :     bot,
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
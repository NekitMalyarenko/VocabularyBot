package telegramHandlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"

	"github.com/NekitMalyarenko/VocabularyBot/db"
	"github.com/NekitMalyarenko/VocabularyBot/telegram/data"
	"github.com/NekitMalyarenko/VocabularyBot/telegram/helpers"
	"github.com/NekitMalyarenko/VocabularyBot/web"
)

const (
	ME = 360952996
)


func NewUserActionHandler(actionData telegramData.ActionData) bool {
	dbManager := db.GetDBManager()

	hasUser, err := dbManager.HasUser(actionData.ChatId)
	if err != nil {
		log.Fatal(err)
		return false
	}

	if !hasUser {
		err = db.GetDBManager().AddUser(db.User{
			Id:        actionData.ChatId,
			FirstName: actionData.Update.Message.Chat.FirstName,
			LastName:  actionData.Update.Message.Chat.LastName,
			UserName:  actionData.Update.Message.Chat.UserName,
			IsTester:  false,
		})
		if err != nil {
			log.Println(err)
			return false
		}

		text := telegramHelpers.MessageBuilderInit().BoldText("Hello! I am the Vocabulary Bot.").NewRow().NormalText(
			"I send you new word with it's definition and examples each day.").NewRow().NormalText("for any "+
			"suggestions contact ").MentionUser(ME, "me").Text

		message := tgbotapi.NewMessage(actionData.ChatId, text)
		message.ParseMode = "HTML"

		actionData.Bot.Send(message)
		return true
	} else {
		message := tgbotapi.NewMessage(actionData.ChatId, "You are already registered!")
		actionData.Bot.Send(message)
		return true
	}
}


func GetWord(actionData telegramData.ActionData) bool {
	word, err := web.GetRandomWord()
	if err != nil {
		log.Println(err)
		return false
	}

	message := tgbotapi.NewMessage(actionData.ChatId, word.ToString())
	message.ParseMode = "HTML"
	actionData.Bot.Send(message)
	return true
}


func Cancel(actionData telegramData.ActionData) bool {
	actionData.Context.ReCreateContext()

	message := tgbotapi.NewMessage(actionData.ChatId, "I agree with you.")
	actionData.Bot.Send(message)
	return true
}
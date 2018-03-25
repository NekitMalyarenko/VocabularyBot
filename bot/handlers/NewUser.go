package handlers

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/NekitMalyarenko/VocabularyBot/types"
	"github.com/NekitMalyarenko/VocabularyBot/db"
	"github.com/NekitMalyarenko/VocabularyBot/telegram"
	"github.com/NekitMalyarenko/VocabularyBot/bot/services"
)

func NewUser(actionData *types.ActionData) bool {
	dbManager := db.GetDBManager()

	hasUser, err := dbManager.HasUser(actionData.ChatId)
	if err != nil {
		log.Fatal(err)
		return false
	}

	if !hasUser {
		err = db.GetDBManager().AddUser(types.User{
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

		text := telegramServices.MessageBuilderInit().BoldText("Hello! I am the Vocabulary Bot.").NewRow().NormalText(
			"I send you new word with it's definition and examples each day.").NewRow().NormalText("for any "+
			"suggestions contact ").MentionUser(telegram.ME, "me").Text

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

package handlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/NekitMalyarenko/VocabularyBot/types"
)


func StatusHandler(actionData *types.ActionData) bool {

	if len(actionData.Context.Data) == 0 {
		actionData.Context.Data["test"] = "hello"
		message := tgbotapi.NewMessage(actionData.ChatId, "ok")
		actionData.Bot.Send(message)
	} else {
		message := tgbotapi.NewMessage(actionData.ChatId, "ok " + actionData.Context.Data["test"].(string))
		actionData.Bot.Send(message)
	}

	return true
}

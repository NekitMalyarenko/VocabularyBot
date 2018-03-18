package telegramHandlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"vocabulary-tg-bot/telegram/data"
)

const (
	START_NN_TRAINIG_BUTTON = 0
)

func BeginNNTrainingButton(actionData telegramData.ActionData) bool {
	actionData.Bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
		ChatID:    actionData.ChatId,
		MessageID: actionData.Update.CallbackQuery.Message.MessageID,
	})
	return BeginNNTraining(actionData)
}

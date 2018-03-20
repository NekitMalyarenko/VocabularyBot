package telegramHandlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
	"log"
	"math"

	"github.com/NekitMalyarenko/VocabularyBot/telegram/data"
	"github.com/NekitMalyarenko/VocabularyBot/telegram/helpers"
	"github.com/NekitMalyarenko/VocabularyBot/web"
	"github.com/NekitMalyarenko/VocabularyBot/web/types"
	"github.com/NekitMalyarenko/VocabularyBot/db"
)


const (
	USER_TRAINING_QUESTIONS = 5
)



func BeginNNTrainingButton(actionData telegramData.ActionData) bool {
	actionData.Bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
		ChatID    : actionData.ChatId,
		MessageID : actionData.Update.CallbackQuery.Message.MessageID,
	})
	return BeginNNTraining(actionData)
}


func BeginNNTraining(actionData telegramData.ActionData) bool {
	actionData.Context.ReCreateContext()
	manager := db.GetDBManager()

	isTester, err := manager.IsUserTester(actionData.ChatId)
	if err != nil {
		log.Println(err)
		return false
	}

	if isTester {
		actionData.Context.NextAction = NNTraining
		actionData.Context.Data["page"] = 2

		word := web.GetNNTrainingWord()
		actionData.Context.Data["word"] = word

		message := tgbotapi.NewMessage(actionData.ChatId, word.ToString())
		message.ParseMode = "HTML"
		actionData.Bot.Send(message)
		return true

	} else {
		text := telegramHelpers.MessageBuilderInit().NormalText("You aren't ").BoldText("Tester").Text
		message := tgbotapi.NewMessage(actionData.ChatId, text)
		message.ParseMode = "HTML"
		actionData.Bot.Send(message)
		return true
	}
}


func NNTraining(actionData telegramData.ActionData) bool {
	var message tgbotapi.MessageConfig

	wordScore, err := strconv.ParseFloat(actionData.Update.Message.Text, 64)
	if err != nil || wordScore > 10.0 || wordScore < 0.0 {
		message = tgbotapi.NewMessage(actionData.ChatId, "Your input was incorrect.Let's try again.\nRate word from 0 - 10")
		actionData.Bot.Send(message)
		return true
	}
	wordScore = toFixed(wordScore/10, 3)

	word := actionData.Context.Data["word"].(*webTypes.RowWordData)
	err = db.GetDBManager().AddNNData(word, wordScore, actionData.ChatId)
	if err != nil {
		log.Println(err)
		return false
	}

	if actionData.Context.Data["page"] == USER_TRAINING_QUESTIONS+1 {
		return EndNNTraining(actionData)
	} else {
		word = web.GetNNTrainingWord()
		actionData.Context.Data["word"] = word

		message = tgbotapi.NewMessage(actionData.ChatId, word.ToString())
		message.ParseMode = "HTML"
		actionData.Context.Data["page"] = actionData.Context.Data["page"].(int) + 1
	}

	actionData.Bot.Send(message)
	return true
}


func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}


func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}


func EndNNTraining(actionData telegramData.ActionData) bool {
	actionData.Context.ReCreateContext()

	text := telegramHelpers.MessageBuilderInit().BoldText("Thank you =)\n").
		NormalText("Your contribution is invaluable.").Text

	message := tgbotapi.NewMessage(actionData.ChatId, text)
	message.ParseMode = "HTML"
	actionData.Bot.Send(message)
	return true
}
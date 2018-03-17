package telegramHandlers

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
	"math"

	"vocabulary-tg-bot/web"
	"vocabulary-tg-bot/web/types"
	"vocabulary-tg-bot/telegram/data"
	"vocabulary-tg-bot/telegram/helpers"
	"vocabulary-tg-bot/db"
)

const (
	ME                      = 360952996
	USER_TRAINING_QUESTIONS = 5
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
			Id        : actionData.ChatId,
			FirstName : actionData.Update.Message.Chat.FirstName,
			LastName  :  actionData.Update.Message.Chat.LastName,
			UserName  :  actionData.Update.Message.Chat.UserName,
			IsTester  :  false,
		})
		if err != nil {
			log.Println(err)
			return false
		}

		text := telegramHelpers.MessageBuilderInit().BoldText("Hello! I am the Vocabulary Bot.").NewRow().NormalText(
			"I send you new word with it's definition and examples each day.").NewRow().NormalText("for any " +
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

//NNTraining - Start
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
	wordScore = toFixed(wordScore / 10, 3)

	log.Println("page", actionData.Context.Data["page"], "score:", wordScore)

	word := actionData.Context.Data["word"].(*webTypes.RowWordData)
	err = db.GetDBManager().AddNNData(word, wordScore, actionData.ChatId)
	if err != nil {
		log.Println(err)
		return false
	}

	if actionData.Context.Data["page"] == USER_TRAINING_QUESTIONS + 1 {
		return EndNNTraining(actionData)
	} else {
		word = web.GetNNTrainingWord()
		actionData.Context.Data["word"] = word

		message = tgbotapi.NewMessage(actionData.ChatId, word.ToString())
		message.ParseMode = "HTML"
		actionData.Context.Data["page"] = actionData.Context.Data["page"].(int) +  1
	}

	actionData.Bot.Send(message)
	return true
}


func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}


func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num * output)) / output
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
//NNTraining - End


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
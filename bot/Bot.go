package bot

import (
	"log"
	"time"
	"github.com/go-telegram-bot-api/telegram-bot-api"

	"github.com/NekitMalyarenko/VocabularyBot/bot/handlers"
	"github.com/NekitMalyarenko/VocabularyBot/telegram"
	"github.com/NekitMalyarenko/VocabularyBot/vars"
	"github.com/NekitMalyarenko/VocabularyBot/db"
	"github.com/NekitMalyarenko/VocabularyBot/bot/services"
)

var (
	telegramBot *telegram.Bot
)


func Start() {
	telegramBot = telegram.Init(vars.GetString(vars.TELEGRAM_BOT_TOKEN), true)

	telegramBot.RegisterTextHandler("/status", handlers.StatusHandler)
	telegramBot.RegisterTextHandler("/start", handlers.NewUser)
	telegramBot.RegisterTextHandler("/begin_test", handlers.BeginNNTraining)

	telegramBot.RegisterButtonHandler(handlers.BeginNNTrainingButton)

	if vars.GetBoolean(vars.BOT_LEARNING) {
		log.Println("BOT LEARNING IS TRUE")
		go startBotLearning()
	}

	msg := tgbotapi.NewMessage(telegram.ME, "Поучимся?)")
	msg.ReplyMarkup = telegramServices.KeyboardBuilderInit().
		NewButton("Окей", false, telegram.GetFuncId(handlers.BeginNNTrainingButton)).
		GetKeyboard()
	telegramBot.Send(msg)

	err := telegramBot.Start()
	if err != nil {
		log.Fatal(err)
	}
}


func startBotLearning() {
	time.Sleep(10 * time.Second)

	var (
		hour, hourTrigger, minute int
	)

	for {
		hourTrigger = vars.GetInt(vars.HOUR_TRIGGER)
		hour        = time.Now().Hour()  + 2
		minute      = time.Now().Minute()

		if hour == hourTrigger && minute == 0 {
			testers, err := db.GetDBManager().GetAllTesters()
			if err != nil {
				log.Println(err)
				panic(err)
			}

			for _, val := range testers {
				msg := tgbotapi.NewMessage(val.Id, "Поучимся?)")
				msg.ReplyMarkup = telegramServices.KeyboardBuilderInit().
					NewButton("Окей", false, telegram.GetFuncId(handlers.BeginNNTrainingButton)).
					GetKeyboard()
				telegramBot.Send(msg)
			}
		}

		if !vars.GetBoolean(vars.BOT_LEARNING) {
			break
		}

		time.Sleep(time.Duration(60 - time.Now().Second()) * time.Second)
	}
}
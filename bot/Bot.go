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
	bot *telegram.Bot
)


func Start() {
	bot = telegram.Init(vars.GetString(vars.TELEGRAM_BOT_TOKEN), true)

	bot.RegisterTextHandler("/status", handlers.StatusHandler)
	bot.RegisterTextHandler("/start", handlers.NewUser)
	bot.RegisterTextHandler("/begin_test", handlers.BeginNNTraining)

	bot.RegisterButtonHandler(handlers.BeginNNTrainingButton)

	if vars.GetBoolean(vars.BOT_LEARNING) {
		log.Println("BOT LEARNING IS TRUE")
		go startBotLearning()
	}

	err := bot.Start()
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
				bot.Send(msg)
			}
		}

		if !vars.GetBoolean(vars.BOT_LEARNING) {
			break
		}

		time.Sleep(time.Duration(60 - time.Now().Second()) * time.Second)
	}
}
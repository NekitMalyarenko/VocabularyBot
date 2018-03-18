package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"

	"github.com/NekitMalyarenko/VocabularyBot/db"
	"github.com/NekitMalyarenko/VocabularyBot/telegram/handlers"
	"github.com/NekitMalyarenko/VocabularyBot/telegram/helpers"
	"github.com/NekitMalyarenko/VocabularyBot/vars"
)


func initBotLearning() {
	time.Sleep(10 * time.Second)

	var (
		hour, hourTrigger, minute int
	)

	for {
		hourTrigger = vars.GetInt(vars.HOUR_TRIGGER)
		hour        = time.Now().Hour()  + 2
		minute      = time.Now().Minute()

		if hour == hourTrigger && minute == minute {
			testers, err := db.GetDBManager().GetAllTesters()
			if err != nil {
				log.Println(err)
				panic(err)
			}

			for _, val := range testers {
				msg := tgbotapi.NewMessage(val.Id, "Поучимся?)")
				msg.ReplyMarkup = telegramHelpers.KeyboardBuilderInit().
					NewButton("Окей", false, telegramHandlers.START_NN_TRAINIG_BUTTON).
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

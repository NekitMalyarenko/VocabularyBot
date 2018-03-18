package telegram

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"time"

	"vocabulary-tg-bot/db"
	"vocabulary-tg-bot/telegram/handlers"
	"vocabulary-tg-bot/telegram/helpers"
	"vocabulary-tg-bot/vars"
)

const HOUR = 20

func initBotLearning() {
	i := 0

	time.Sleep(10 * time.Second)

	for {
		hour := time.Now().Hour()
		minute := time.Now().Minute()

		log.Println("Hour", hour, "minute", minute)

		if i == 0 {
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
		time.Sleep(60 * time.Second)
		i++
	}
}
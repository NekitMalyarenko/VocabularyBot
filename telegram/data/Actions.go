package telegramData

import (
	"log"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)


type ActionData struct {
	Update  tgbotapi.Update
	Context *Context
	ChatId  int64
	Bot     *tgbotapi.BotAPI
}

type ActionsHolder struct {
	actions map[string]func(data ActionData)bool
}

var (
	actionsHolder *ActionsHolder
)


func GetActionsHolder() *ActionsHolder {
	if actionsHolder == nil {
		actionsHolder = &ActionsHolder{
			actions: make(map[string]func(data ActionData)bool, 0),
		}
		log.Println("New Instance of Actions holder")
	}

	return actionsHolder
}


func (holder *ActionsHolder) AddAction(key string, action func(data ActionData)bool) {
	holder.actions[key] = action
}


func (holder *ActionsHolder) GetAction(key string) func(data ActionData)bool {
	return holder.actions[key]
}
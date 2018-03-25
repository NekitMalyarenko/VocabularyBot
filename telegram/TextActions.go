package telegram

import (
	"log"
	"github.com/NekitMalyarenko/VocabularyBot/types"
)


type ActionsHolder struct {
	actions map[string]func(data *types.ActionData) bool
}


var (
	actionsHolder *ActionsHolder
)


func getTextActionsHolder() *ActionsHolder {
	if actionsHolder == nil {
		actionsHolder = &ActionsHolder{
			actions: make(map[string]func(data *types.ActionData) bool, 0),
		}
		log.Println("New Instance of Actions holder")
	}

	return actionsHolder
}


func (holder *ActionsHolder) registerAction(key string, action func(data *types.ActionData) bool) {
	holder.actions[key] = action
}


func (holder *ActionsHolder) getAction(key string) func(data *types.ActionData) bool {
	return holder.actions[key]
}

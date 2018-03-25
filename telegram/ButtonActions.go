package telegram

import (
	"log"
	"runtime"
	"reflect"
	"github.com/NekitMalyarenko/VocabularyBot/types"
)


type ButtonsHolder struct {
	data map[string]func(*types.ActionData)bool
}


var (
	buttonsHolder *ButtonsHolder
)


func getButtonActionsHolder() *ButtonsHolder {

	if buttonsHolder == nil {
		buttonsHolder = &ButtonsHolder{
			data: make(map[string]func(*types.ActionData)bool, 0),
		}
		log.Println("New Instance of ButtonsData holder")
	}

	return buttonsHolder
}


func (holder *ButtonsHolder) registerAction(function func(*types.ActionData)bool) {
	funcName := runtime.FuncForPC(reflect.ValueOf(function).Pointer()).Name()
	holder.data[funcName] = function
}


func (holder *ButtonsHolder) getAction(id string) func(*types.ActionData)bool {
	return holder.data[id]
}

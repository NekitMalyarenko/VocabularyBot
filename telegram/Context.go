package telegram

import (
	"log"
	"github.com/NekitMalyarenko/VocabularyBot/types"
)


type ContextHolder struct {
	data map[int64]*types.Context
}


var (
	contextHolder *ContextHolder
)


func getContextHolder() *ContextHolder {

	if contextHolder == nil {
		contextHolder = &ContextHolder{
			data: make(map[int64]*types.Context, 0),
		}
		log.Println("New Instance of Context holder")
	}

	return contextHolder
}


func (holder *ContextHolder) createContext(key int64) {
	holder.data[key] = new(types.Context)
	holder.data[key].Data = make(map[string]interface{}, 0)
	log.Println("Creating Context for", key)
}


func (holder *ContextHolder) hasContext(key int64) bool {
	return holder.data[key] != nil
}


func (holder *ContextHolder) getContext(key int64) *types.Context {
	return holder.data[key]
}
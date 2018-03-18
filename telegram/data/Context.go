package telegramData

import (
	"log"
)

type Context struct {
	ChatId     int64
	Data       map[string]interface{}
	NextAction func(actionData ActionData) bool
}

type ContextHolder struct {
	data map[int64]*Context
}

var (
	contextHolder *ContextHolder
)

func GetContextHolder() *ContextHolder {

	if contextHolder == nil {
		contextHolder = &ContextHolder{
			data: make(map[int64]*Context, 0),
		}
		log.Println("New Instance of Context holder")
	}

	return contextHolder
}

func (holder *ContextHolder) CreateContext(key int64) {
	holder.data[key] = new(Context)
	holder.data[key].Data = make(map[string]interface{}, 0)
	log.Println("Creating Context for", key)
}

func (holder *ContextHolder) HasContext(key int64) bool {
	return holder.data[key] != nil
}

func (holder *ContextHolder) GetContext(key int64) *Context {
	return holder.data[key]
}

func (context *Context) ReCreateContext() {
	context.Data = make(map[string]interface{}, 0)
	context.NextAction = nil
}

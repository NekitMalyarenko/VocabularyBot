package types

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type NNData struct {
	Id            int64   `db:"id,omitempty"`
	Word          string  `db:"word"`
	WordRank      int     `db:"word_rank"`
	Definitions   string  `db:"definitions"`
	UsageExamples string  `db:"usage_examples"`
	WordRating    float64 `db:"word_rating"`
	RatedUserId   int64   `db:"rated_by"`
}


type User struct {
	Id        int64  `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	UserName  string `db:"user_name"`
	IsTester  bool   `db:"is_tester"`
}


type Word struct {
	Id            int64  `db:"id"`
	Word          string `db:"word"`
	Date          string `db:"date"`
	Definitions   string `db:"definitions"`
	UsageExamples string `db:"usage_examples"`
	Likes         int    `db:"likes"`
	Dislikes      int    `db:"dislikes"`
}


type RowWordData struct {
	Word          string
	Definitions   map[string][]string
	UsageExamples map[string][]string
	WordRank      int
}


type ActionData struct {
	Update     tgbotapi.Update
	Bot        *tgbotapi.BotAPI
	Context    *Context
	ButtonData map[string]interface{}
	ChatId     int64
}


type Context struct {
	ChatId     int64
	Data       map[string]interface{}
	NextAction func(actionData *ActionData) bool
}

func (context *Context) Clear() {
	context.Data = make(map[string]interface{}, 0)
	context.NextAction = nil
}
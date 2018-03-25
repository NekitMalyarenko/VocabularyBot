package main

import (
	"github.com/NekitMalyarenko/VocabularyBot/db"
	"github.com/NekitMalyarenko/VocabularyBot/bot"
)


const (
	ME = 360952996
)


func main() {
	db.GetDBManager()
	defer db.CloseConnection()

	bot.Start()
}

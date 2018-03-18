package main

import (
	"github.com/NekitMalyarenko/VocabularyBot/db"
	"github.com/NekitMalyarenko/VocabularyBot/telegram"
)

func main() {
	db.GetDBManager()
	defer db.CloseConnection()

	telegram.Start()
}
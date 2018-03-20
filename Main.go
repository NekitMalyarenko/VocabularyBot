package main

import (
	"github.com/NekitMalyarenko/VocabularyBot/telegram"
	"github.com/NekitMalyarenko/VocabularyBot/db"
)

func main() {
	db.GetDBManager()
	defer db.CloseConnection()

	telegram.Start()
}
package main

import (
	"os"
	"vocabulary-tg-bot/telegram"
	"vocabulary-tg-bot/vars"
	"vocabulary-tg-bot/db"
)

func main() {
	db.GetDBManager()
	defer db.CloseConnection()

	telegram.Start()
}

package main

import (
	"os"
	"github.com/NekitMalyarenko/VocabularyBot/vars"
	"github.com/NekitMalyarenko/VocabularyBot/db"
	"github.com/NekitMalyarenko/VocabularyBot/bot"
)


const (
	ME = 360952996
)


func main() {
	os.Setenv(vars.DB_CONNECTION_STRING, "user=uunrcwvmqzwoap password=eb0e8547c23a5a560e54021dca704a896dfad0c039b187d1a32e8f5b5038c37a host=ec2-54-247-95-125.eu-west-1.compute.amazonaws.com port=5432 database=d78hor7l3bje7q sslmode=require")
	os.Setenv(vars.BOT_LEARNING, "false")
	os.Setenv(vars.TELEGRAM_BOT_TOKEN, "522818795:AAFQnTgc-nfziv3zXjb7MNF1PzoSSIjanHI")

	db.GetDBManager()
	defer db.CloseConnection()

	bot.Start()
}

package main

import (
	"os"
	"vocabulary-tg-bot/vars"
	"vocabulary-tg-bot/db"
	"vocabulary-tg-bot/telegram"
)




func main() {
	os.Setenv(vars.TELEGRAM_BOT_TOKEN, "554232117:AAEZo-oWiMzQ97egh-q21_dZ_AFnHIuokeE")
	os.Setenv(vars.BOT_LEARNING, "true")
	os.Setenv(vars.DB_CONNECTION_STRING, "user=uunrcwvmqzwoap password=eb0e8547c23a5a560e54021dca704a896dfad0c039b187d1a32e8f5b5038c37a host=ec2-54-247-95-125.eu-west-1.compute.amazonaws.com port=5432 database=d78hor7l3bje7q sslmode=require")

	db.GetDBManager()
	defer db.CloseConnection()

	telegram.Start()
}
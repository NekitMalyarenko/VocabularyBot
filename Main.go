package main

import (
	"os"
	"vars"
	"telegram"
	"db"
)




func main() {
	os.Setenv(vars.TELEGRAM_BOT_TOKEN, "554232117:AAEZo-oWiMzQ97egh-q21_dZ_AFnHIuokeE")
	os.Setenv(vars.DB_CONNECTION_STRING, "user=uunrcwvmqzwoap password=eb0e8547c23a5a560e54021dca704a896dfad0c039b187d1a32e8f5b5038c37a host=ec2-54-247-95-125.eu-west-1.compute.amazonaws.com port=5432 database=d78hor7l3bje7q sslmode=require")

	db.GetDBManager()
	defer db.CloseConnection()

	telegram.Start()

	/*word := web.RowWordData{
		Word:"test",
		WordRank:0,
		UsageExamples:[]string{"ex", "ex2", "ex3"},
		Definitions:[]string{"def", "def2", "def3"},

	}*/

	/*word, err := web.GetRandomWord()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(word.Word)
	log.Println(word.WordRank)
	log.Println(word.Definitions)
	log.Println(word.UsageExamples)*/
}
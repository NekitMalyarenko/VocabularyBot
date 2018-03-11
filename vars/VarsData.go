package vars

import (
	"os"
	"strconv"
	"log"
)

const (
	TELEGRAM_BOT_TOKEN  = "BOT_TOKEN"
	DB_CONNECTION_STRING = "test"
	DB_HOST             = "DB_HOST"
	DB_NAME             = "DB_NAME"
	DB_USER             = "DB_USER"
	DB_PASSWORD         = "DB_PASSWORD"
	DB_SOCKET           = "DB_SOCKET"
)


func GetString(varName string) string {
	return os.Getenv(varName)
}


func GetInt(varName string) int {
	temp, err := strconv.Atoi(os.Getenv(varName))
	if err != nil {
		log.Fatal(err)
		return -1
	} else {
		return temp
	}
}

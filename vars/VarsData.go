package vars

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	TELEGRAM_BOT_TOKEN   = "BOT_TOKEN"
	DB_CONNECTION_STRING = "DB_CONNECTION_STRING"
	BOT_LEARNING         = "BOT_LEARNING"
	HOUR_TRIGGER         = "HOUR_TRIGGER"
)

func GetString(varName string) string {
	return os.Getenv(varName)
}

func GetBoolean(varName string) bool {
	return strings.ToLower(os.Getenv(varName)) == "true"
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

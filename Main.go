package main

import (
	"telegram"
	"db"
)

func main() {
	db.GetDBManager()
	defer db.CloseConnection()

	telegram.Start()
}

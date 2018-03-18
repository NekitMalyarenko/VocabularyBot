package db

import (
	"log"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/postgresql"
	"github.com/NekitMalyarenko/VocabularyBot/vars"
)

type dbManager struct {
	Session sqlbuilder.Database
}

var (
	manager *dbManager
)

func GetDBManager() *dbManager {

	if manager == nil {
		connection, err := postgresql.ParseURL(vars.GetString(vars.DB_CONNECTION_STRING))
		if err != nil {
			return nil
		}
		//connection.Options = options

		log.Println(connection.Options)

		session, err := postgresql.Open(connection)
		if err != nil {
			return nil
		}

		log.Println("Creating New DB Connection")

		manager = &dbManager{Session: session}
	}

	return manager
}

func CloseConnection() {
	manager.Session.Close()
	log.Println("DB connection was closed")
}

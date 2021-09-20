package main

import (
	"github.com/jim-nnamdi/go_crud/database"
)

func main() {
	config := database.Config{
		Servername: "localhost:3306",
		User:       "root",
		Password:   "root",
		DB:         "wickedgo",
	}

	connectionString := database.getConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
}

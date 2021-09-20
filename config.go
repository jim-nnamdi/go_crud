package main

import "fmt"

type Config struct {
	Servername string
	User       string
	Password   string
	DB         string
}

var getConnectionString = func(config Config) string {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.User, config.Password, config.Servername, config.DB)

	return connectionString
}

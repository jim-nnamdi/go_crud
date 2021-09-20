package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

func connect(connectionString string) error {
	var err error
	Connector, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("connection was successful")
	return nil
}

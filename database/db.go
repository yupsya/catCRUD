package database

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnToDB() {
	var err error
	dsn := "host=localhost user=postgres password=1234 dbname=dotnetgo port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to pg database: ", err)
	}
}
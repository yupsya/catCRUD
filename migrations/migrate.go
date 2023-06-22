package main

import (
	"go-gin-gorm/database"
	"go-gin-gorm/models"
)

func init() {
	database.ConnToDB()
}

func main() {
	database.DB.AutoMigrate(&models.Cat{})
}
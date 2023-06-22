package controllers

import (
	"go-gin-gorm/database"
	"go-gin-gorm/models"

	"github.com/gin-gonic/gin"
)

func CatCreate (c *gin.Context) {

	var body struct {
		Owner 		string
		Name 		string 
		Breed       string 
		DateOfBirth string 
	}

	c.Bind(&body)

	cat := models.Cat{Owner: body.Owner, Name: body.Name, Breed: body.Breed}
	res := database.DB.Create(&cat)

	if res.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"cat": cat}) // gin.H is handler (type map[string]any)
}

func CatsGet(c *gin.Context) {
	var cats []models.Cat
	database.DB.Find(&cats)

	c.JSON(200, gin.H{"cats": cats})
}

func CatRetrieve(c *gin.Context) {
	var cat models.Cat
	database.DB.First(&cat, c.Param("id"))

	c.JSON(200, gin.H{"cat": cat})
}

func CatUpdate(c *gin.Context) {
	// нашли китика
	var cat models.Cat
	database.DB.First(&cat, c.Param("id"))

	// объявили структуру для китика
	var body struct {
		Owner 		string
		Name 		string 
		Breed       string 
		DateOfBirth string 
	}
	
	// апдейтим китика данными из структуры
	database.DB.Model(&cat).Updates(models.Cat{
		Owner: body.Owner,
		Name: body.Name,
		Breed: body.Breed,
		DateOfBirth: body.DateOfBirth,
	})
	c.JSON(200, gin.H{"cat": cat})
}

func CatDelete(c *gin.Context) {
	
	// удалили китика
	database.DB.Delete(&models.Cat{}, c.Param("id"))
	c.Status(200)
}
package main

import (
	"go-gin-gorm/controllers"
	"go-gin-gorm/database"
	"github.com/gin-gonic/gin"
)

func init() {
	database.ConnToDB()
}

func main() {
  r := gin.Default()

  r.POST("/create", controllers.CatCreate)
  r.GET("/cats", controllers.CatsGet)
  r.GET("/cats/:id", controllers.CatRetrieve)
  r.PUT("/cats/:id", controllers.CatUpdate)
  r.DELETE("/cats/:id", controllers.CatDelete)

  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
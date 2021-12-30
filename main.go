package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raublekick/go-api-example/controllers"
	"github.com/raublekick/go-api-example/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// TODO: how to simplify config for many endpoints? Group routing!
	r.GET("/games", controllers.FindGames)
	r.GET("/games/:id", controllers.FindGame)
	r.POST("/games", controllers.CreateGame)
	r.PATCH("/games/:id", controllers.UpdateGame)
	r.DELETE("/games/:id", controllers.DeleteGame)

	r.Run()
}

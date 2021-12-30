package main

import (
	"github.com/gin-gonic/gin"
	"github.com/raublekick/go-api-example/controllers"
	"github.com/raublekick/go-api-example/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	// TODO: how to simplify config for many endpoints?
	r.GET("/games", controllers.FindGames)

	r.Run()
}

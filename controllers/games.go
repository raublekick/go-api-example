// controllers/games.go

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raublekick/go-api-example/models"
)

// GET /games
// Get all games
func FindGames(c *gin.Context) {
	var games []models.Game
	models.DB.Find(&games)

	c.JSON(http.StatusOK, gin.H{"data": games})
}

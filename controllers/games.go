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

// GET /games/:id
// Get a single game
func FindGame(c *gin.Context) {
	var game models.Game
	if err := models.DB.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// POST /games
// Create new game
func CreateGame(c *gin.Context) {
	var input models.CreateGameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	game := models.Game{Title: input.Title, URL: input.URL}
	models.DB.Create(&game)

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// PATCH /games/:id
// Update game
func UpdateGame(c *gin.Context) {
	var game models.Game

	// validate record exists
	if err := models.DB.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate input
	var input models.UpdateGameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&game).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": game})
}

// DELETE /games/:id
// Delete game
func DeleteGame(c *gin.Context) {
	var game models.Game

	// validate record exists
	if err := models.DB.Where("id = ?", c.Param("id")).First(&game).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&game)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

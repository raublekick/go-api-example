// models.game.go

package models

type Game struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

type CreateGameInput struct {
	Title string `json:"title" binding:"required"`
	URL   string `json:"url" binding:"required"`
}

type UpdateGameInput struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

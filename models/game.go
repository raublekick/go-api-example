// models.game.go

package models

type Game struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

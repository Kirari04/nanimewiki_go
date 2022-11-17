package models

type Anime struct {
	ID    uint   `json:"id" gorm:"primary_key" binding:"required"`
	Title string `json:"title" binding:"required"`
}

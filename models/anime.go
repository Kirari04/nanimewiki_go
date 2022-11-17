package models

type Anime struct {
	ID    uint   `json:"id" gorm:"primary_key;autoIncrement;not null;unique"`
	Title string `json:"title"`
}

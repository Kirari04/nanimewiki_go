package models

type Anime struct {
	DBModel
	Title string `json:"title"`
	Type string `json:"type"`
	Status string `json:"status"`
}

package models

type User struct {
	DBModel
	Email         string `gorm:"not null;uniqueIndex"`
	EmailVerified bool   `gorm:"default:false"`
	Username      string `gorm:"not null;index"`
	Password      string `gorm:""`
	Avatar        string `gorm:""`
}

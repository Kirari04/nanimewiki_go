package models

type User struct {
	ID       uint   `gorm:"primary_key;autoIncrement;not null;unique"`
	Email    string `gorm:"not null;uniqueIndex"`
	Username string `gorm:"not null;index"`
	Password string `gorm:""`
	Avatar   string `gorm:""`
	Updated  int64  `gorm:"autoUpdateTime"` // Use unix milli seconds as updating time
	Created  int64  `gorm:"autoCreateTime"`
}

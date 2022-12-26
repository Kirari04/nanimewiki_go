package models

import "time"

type User struct {
	DBModel
	Email         string `gorm:"not null;uniqueIndex;size:256"`
	EmailVerified bool   `gorm:"default:false"`
	Username      string `gorm:"not null;index;size:256"`
	Password      string `gorm:""`
	Avatar        string `gorm:""`
}

type EmailVerificationCode struct {
	DBModel
	UserID uint      `gorm:"not null"`
	User   User      `gorm:""`
	Code   string    `gorm:""`
	Expire time.Time `gorm:"not null"`
}

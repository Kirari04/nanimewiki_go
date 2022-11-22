package models

import "time"

type User struct {
	DBModel
	Email         string `gorm:"not null;uniqueIndex"`
	EmailVerified bool   `gorm:"default:false"`
	Username      string `gorm:"not null;index"`
	Password      string `gorm:""`
	Avatar        string `gorm:""`
}

type EmailVerificationKey struct {
	DBModel
	UserID uint   `gorm:""`
	Key    string `gorm:""`
	Expire time.Time
}

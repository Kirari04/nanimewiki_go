package models

import (
	"time"

	"gorm.io/gorm"
)

type DBModel struct {
	ID        uint `json:"id" gorm:"primary_key;autoIncrement;not null;unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

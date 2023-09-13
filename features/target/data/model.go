package data

import (
	"time"

	"gorm.io/gorm"
)

type Target struct {
	ID           		string `gorm:"primaryKey"`
	CreatedAt    		time.Time
	UpdatedAt    		time.Time
	DeletedAt    		gorm.DeletedAt `gorm:"index"`
	KontenTarget 		string
	Status       		string
	DevisiID     		string
	UserIDPembuat       string
}
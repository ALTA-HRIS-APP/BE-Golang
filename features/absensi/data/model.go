package data

import (
	"time"

	"gorm.io/gorm"
) 

type Absensi struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    string
	OverTime  time.Time
	JamMasuk time.Time
	JamKeluar time.Time
}
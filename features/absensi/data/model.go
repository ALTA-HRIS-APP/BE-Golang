package data

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
) 

type Absensi struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    string
	OverTime  time.Time
	JamMasuk time.Time
	JamKeluar time.Time
}
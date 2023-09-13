package data

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Target struct {
	ID           		uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	CreatedAt    		time.Time
	UpdatedAt    		time.Time
	DeletedAt    		gorm.DeletedAt `gorm:"index"`
	KontenTarget 		string
	Status       		string
	DevisiID     		string
	UserIDPembuat       string
}
package data

import (
	"time"

	"gorm.io/gorm"
)

type Target struct {
	ID             string `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	KontenTarget   string         `gorm:"column:konten_target;not null"`
	Status         string         `gorm:"type:enum('not completed','completed');column:status;"`
	DevisiID       string         `gorm:"column:devisi_id;not null"`
	UserIDPembuat  string         `gorm:"column:user_id_pembuat;not null"`
	UserIDPenerima string         `gorm:"column:user_id_penerima;not null"`
	Due_Date       string         `gorm:"column:due_date;not null"`
	Proofs         []string       `gorm:"column:proofs"`
	// User           User           `gorm:"foreignKey:UserIDPembuat"`
}

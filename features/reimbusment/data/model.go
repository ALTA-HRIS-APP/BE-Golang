package data

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reimbursement struct {
	ID        		string `gorm:"primaryKey"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	DeletedAt 		gorm.DeletedAt `gorm:"index"`
	Description 	string
	Status 			string
	BatasanReimburs int
	Nominal 		int
	Tipe 			string
	Persetujuan 	string
	UrlBukti 		string
	UserID 			uuid.UUID

}
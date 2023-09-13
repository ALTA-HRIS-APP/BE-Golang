package data

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Reimbursement struct {
	ID        		uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
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
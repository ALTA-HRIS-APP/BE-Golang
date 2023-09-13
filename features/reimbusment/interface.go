package reimbusment

import (
	"time"
)

type ReimbursementEntity struct {
	ID              string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
	Description     string
	Status          string
	BatasanReimburs int
	Nominal         int
	Tipe            string
	Persetujuan     string
	UrlBukti        string
	UserID          string
}
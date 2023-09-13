package target

import (
	"time"
)

type TargetEntity struct {
	ID            string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
	KontenTarget  string
	Status        string
	DevisiID      string
	UserIDPembuat string
}
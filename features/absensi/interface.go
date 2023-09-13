package absensi

import (
	"time"
)

type AbsensiEntity struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	UserID    string
	OverTime  time.Time
	JamMasuk  time.Time
	JamKeluar time.Time
}
package cuti

import (
	"time"
)

type CutiEntity struct {
	ID           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	TipeCuti     string
	Status       string
	JumlahCuti   int
	BatasanCuti  int
	Description  string
	Persetujuan  time.Time
	StartCuti    time.Time
	EndCuti      time.Time
	UrlPendukung string
	UserID       string
}
package cuti

import (
	"time"
)

type CutiEntity struct {
	ID           string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
	TipeCuti     string `validate:"required"`
	Status       string
	JumlahCuti   int `validate:"required"`
	BatasanCuti  int
	Description  string `validate:"required"`
	Persetujuan  string
	StartCuti    string
	EndCuti      string
	UrlPendukung string `validate:"required"`
	UserID       string
}

type CutiDataInterface interface {
	Insert(input CutiEntity) error
}
type CutiServiceInterface interface {
	Add(input CutiEntity) error
}

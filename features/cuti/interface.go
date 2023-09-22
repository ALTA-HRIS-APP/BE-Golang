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
	User         UserEntity
}

type UserEntity struct {
	ID   string
	Name string
}

type CutiDataInterface interface {
	Insert(input CutiEntity) error
	SelectAllKaryawan(idUser string) ([]CutiEntity, error)
	SelectAll(token string,) ([]CutiEntity, error)
	SelectById(id string) (CutiEntity, error)
	UpdateKaryawan(input CutiEntity, id string) error
	Update(input CutiEntity, id string) error
	Delete(id string) error
}
type CutiServiceInterface interface {
	Add(input CutiEntity) error
	Get(token string,idUser string) ([]CutiEntity, error)
	Edit(input CutiEntity, id string, idUser string) error
	Delete(id string) error
}

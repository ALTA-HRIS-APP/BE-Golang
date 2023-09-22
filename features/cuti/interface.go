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
type PenggunaEntity struct {
	ID          string
	NamaLengkap string
	Jabatan     string
	Devisi      string
}

type QueryParams struct {
	Page             int
	ItemsPerPage     int
	SearchName       string
	IsClassDashboard bool
}

type CutiDataInterface interface {
	Insert(input CutiEntity) error

	SelectAllKaryawan(idUser string, param QueryParams) (int64, []CutiEntity, error)
	SelectAll(token string, param QueryParams) (int64, []CutiEntity, error)

	SelectById(id string) (CutiEntity, error)
	UpdateKaryawan(input CutiEntity, id string) error
	Update(input CutiEntity, id string) error
	Delete(id string) error
	SelectUserById(idUser string) (PenggunaEntity, error)
}
type CutiServiceInterface interface {
	Add(input CutiEntity) error

	Get(token string, idUser string, param QueryParams) (bool, []CutiEntity, error)

	Edit(input CutiEntity, id string, idUser string) error
	Delete(id string) error
	GetCutiById(id string) (CutiEntity, error)
}

package absensi

import (
	"time"
)

type AbsensiEntity struct {
	ID              string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
	UserID          string
	OverTimeMasuk   string
	OverTimePulang  string
	JamMasuk        string
	JamKeluar       string
	TanggalSekarang string
	User            UserEntity
}

type UserEntity struct {
	ID   string
	Name string
}

type QueryParams struct {
	Page             int
	ItemsPerPage     int
	SearchName       string
	IsClassDashboard bool
}

type AbsensiDataInterface interface {
	SelectAllKaryawan(idUser string, param QueryParams) (int64, []AbsensiEntity, error)
	Insert(input AbsensiEntity) error
	Update(input AbsensiEntity, idUser string, id string) error
	SelectById(id string) (AbsensiEntity, error)
	SelectAll(param QueryParams) (int64, []AbsensiEntity, error)
}

type AbsensiServiceInterface interface {
	Get(idUser string, param QueryParams) (bool, []AbsensiEntity, error)
	Add(idUser string) error
	Edit(idUser string, id string) error
	SelectById(id string) (AbsensiEntity, error)
}

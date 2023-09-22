package target

import (
	"time"
)

type TargetEntity struct {
	ID             string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
	KontenTarget   string `validate:"required"`
	Status         string
	DevisiID       string `validate:"required"`
	UserIDPembuat  string
	UserIDPenerima string `validate:"required"`
	DueDate        string `validate:"required"`
	Proofs         string
	User           UserEntity
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
type QueryParam struct {
	Page           int
	LimitPerPage   int
	SearchKonten   string
	SearchStatus   string
	ExistOtherPage bool
}
type TargetDataInterface interface {
	Insert(input TargetEntity) (string, error)
	SelectAll(token string, param QueryParam) (int64, []TargetEntity, error)
	Select(idTarget string) (TargetEntity, error)
	Update(idTarget string, targetData TargetEntity) error
	Delete(idTarget string) error
	GetUserByIDAPI(idUser string) (PenggunaEntity, error)
	SelectAllKaryawan(idUser string, param QueryParam) (int64, []TargetEntity, error)
}

type TargetServiceInterface interface {
	Create(input TargetEntity) (string, error)
	GetAll(token string, idUser string, param QueryParam) (bool, []TargetEntity, error)
	GetById(idTarget string, idUser string) (TargetEntity, error)
	UpdateById(idTarget string, idUser string, targetData TargetEntity) error
	DeleteById(idTarget string, idUser string) error
}

package target

import (
	apinodejs "be_golang/klp3/features/apiNodejs"
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
	Due_Date       string `validate:"required"`
	Proofs         string
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
	SelectAll(param QueryParam) (int64, []TargetEntity, error)
	Select(targetID string) (TargetEntity, error)
	Update(targetID string, targetData TargetEntity) error
	Delete(targetID string) error
	GetUserByIDAPI(idUser string) (apinodejs.Pengguna, error)
}

type TargetServiceInterface interface {
	Create(input TargetEntity) (string, error)
	GetAll(userID string, param QueryParam) (bool, []TargetEntity, error)
	GetById(targetID string, userID string) (TargetEntity, error)
	UpdateById(targetID string, userID string, targetData TargetEntity) error
	DeleteById(targetID string, userID string) error
	GetUserByIDAPI(idUser string) (apinodejs.Pengguna, error)
}

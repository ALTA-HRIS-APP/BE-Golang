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
	Due_Date       string `validate:"required"`
	Proofs         string
}

type TargetDataInterface interface {
	Insert(input TargetEntity) (string, error)
	SelectAll(userID string) ([]TargetEntity, error)
	Select(targetID string, userID string) (TargetEntity, error)
	Update(targetID string, userID string, targetData TargetEntity) error
	Delete(targetID string, userID string) error
}

type TargetServiceInterface interface {
	Create(input TargetEntity) (string, error)
	GetAll(userID string) ([]TargetEntity, error)
	GetById(targetID string, userID string) (TargetEntity, error)
	UpdateById(targetID string, userID string, targetData TargetEntity) error
	DeleteById(targetID string, userID string) error
}

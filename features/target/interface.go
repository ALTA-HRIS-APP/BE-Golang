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
	// User           UserEntity
}

type TargetDataInterface interface {
	Insert(input TargetEntity) (string, error)
	// SelectAll(userID uint) ([]CoreTarget, error)
	// Select(projectId uint, userID uint) (CoreTarget, error)
	// Update(projectId uint, userID uint, projectData CoreTarget) error
	// Delete(projectId uint, userID uint) error
}

type TargetServiceInterface interface {
	Create(input TargetEntity) error
	// GetAll(userID uint) ([]CoreTarget, error)
	// GetById(userID uint, projectId uint) (CoreTarget, error)
	// UpdateById(projectId uint, userID uint, projectData CoreTarget) error
	// DeleteById(projectId uint, userID uint) error
}

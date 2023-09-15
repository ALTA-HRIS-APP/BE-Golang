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
	UserIDPembuat  string `validate:"required"`
	UserIDPenerima string `validate:"required"`
	Due_Date       string `validate:"required"`
	Proofs         []string
	// User           UserEntity
}

type ProjectDataInterface interface {
	Insert(input TargetEntity) (string, error)
	// SelectAll(userID uint) ([]CoreProject, error)
	// Select(projectId uint, userID uint) (CoreProject, error)
	// Update(projectId uint, userID uint, projectData CoreProject) error
	// Delete(projectId uint, userID uint) error
}

type ProjectServiceInterface interface {
	Create(input TargetEntity) (string, error)
	// GetAll(userID uint) ([]CoreProject, error)
	// GetById(userID uint, projectId uint) (CoreProject, error)
	// UpdateById(projectId uint, userID uint, projectData CoreProject) error
	// DeleteById(projectId uint, userID uint) error
}

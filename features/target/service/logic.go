package service

import (
	"be_golang/klp3/features/target"
	"errors"
	"log"

	"github.com/go-playground/validator/v10"
)

type targetService struct {
	targetRepo target.TargetDataInterface
	validate   *validator.Validate
}

func New(repo target.TargetDataInterface) target.TargetServiceInterface {
	return &targetService{
		targetRepo: repo,
		validate:   validator.New(),
	}
}

// Create implements target.TargetServiceInterface.
func (s *targetService) Create(input target.TargetEntity) (string, error) {
	// responseUser, err := usernodejs.GetByIdUser(input.UserIDPembuat)
	// if err != nil {
	// 	log.Printf("Error get detail user: %s", err.Error())
	// 	return "",err
	// }
	err := s.validate.Struct(input)
	if err != nil {
		log.Printf("Error validate: %s", err.Error())
		return "", errors.New("error validate, konten target, due date required")
	}

	// if responseUser.Jabatan=="manager"

	targetID, err := s.targetRepo.Insert(input)
	if err != nil {
		log.Printf("Error creating target: %s", err.Error())
		return "", err
	}
	log.Println("Target created successfully")
	return targetID, nil
}

// GetAll implements target.TargetServiceInterface.
func (s *targetService) GetAll(userID string) ([]target.TargetEntity, error) {
	result, err := s.targetRepo.SelectAll(userID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetById implements target.TargetServiceInterface.
func (s *targetService) GetById(targetID string, userID string) (target.TargetEntity, error) {
	result, err := s.targetRepo.Select(targetID, userID)
	if err != nil {
		return target.TargetEntity{}, err
	}
	return result, nil
}

// UpdateById implements target.TargetServiceInterface.
func (s *targetService) UpdateById(targetID string, userID string, targetData target.TargetEntity) error {
	err := s.targetRepo.Update(targetID, userID, targetData)
	if err != nil {
		return err
	}
	return nil
}

// DeleteById implements target.TargetServiceInterface.
func (s *targetService) DeleteById(targetID string, userID string) error {
	err := s.targetRepo.Delete(targetID, userID)
	if err != nil {
		return err
	}
	return nil
}
